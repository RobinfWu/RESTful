package main

// fmt - formatted I/O
// log - logging package
// net/http - provides HTTP client and server implementations
// encoding/json - encoding with json format
// mux - a request router and dispatcher for matching incoming requests to their respective handler
// jwt - JSON Web Tokens for Authentication

import (
    "fmt"
    "log"
    "net/http"
    "encoding/json"
    "github.com/gorilla/mux"
    jwt "github.com/dgrijalva/jwt-go"
)

var mySigningKey = []byte("secretkey")

// A homepage that reveals the messeage for authorized clients
func homePage(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Only Authorized Clients can see this message!")
    fmt.Println("Endpoint Hit: homePage")
}

// Validates the token of incoming requests. 
func isAuthorized(endpoint func(http.ResponseWriter, *http.Request)) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

        if r.Header["Token"] != nil {

            token, err := jwt.Parse(r.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
                if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
                    return nil, fmt.Errorf("There was an error")
                }
                return mySigningKey, nil
            })

            if err != nil {
                fmt.Fprintf(w, err.Error())
            }

            if token.Valid {
                endpoint(w, r)
            }
        } else {

            fmt.Fprintf(w, "Not Authorized")
        }
    })
}

// Patient struct (Model)
type Patient struct {
    ID     string  `json:"id"`
    Firstname   string  `json:"firstname"`
    Lastname  string  `json:"lastname"`
    Address  string  `json:"address"`
    Doctor *Doctor `json:"doctor"`
}

// Doctor struct
type Doctor struct {
    Firstname string `json:"firstname"`
    Lastname  string `json:"lastname"`
}

// Initize Patient var as a slice Patient struct
var Patients []Patient

// Get all Patients
func getPatients(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(Patients)
}

// Get single Patient
func getPatient(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    params := mux.Vars(r) // Gets params
    // Loop through patients and find one with the id from the params
    for _, item := range Patients {
        if item.ID == params["id"] {
            json.NewEncoder(w).Encode(item)
            return
        }
    }
    json.NewEncoder(w).Encode(&Patient{})
}

// Add a new Patient
func addPatient(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    var patient Patient
    _ = json.NewDecoder(r.Body).Decode(&patient)
    Patients = append(Patients, patient)
    json.NewEncoder(w).Encode(patient)
}

// Update Patient
func updatePatient(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    params := mux.Vars(r)
    for index, item := range Patients {
        if item.ID == params["id"] {
            Patients = append(Patients[:index], Patients[index+1:]...)
            var patient Patient
            _ = json.NewDecoder(r.Body).Decode(&patient)
            patient.ID = params["id"]
            Patients = append(Patients, patient)
            json.NewEncoder(w).Encode(patient)
            return
        }
    }
}

// Delete a Patient
func deletePatient(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    params := mux.Vars(r)
    for index, item := range Patients {
        if item.ID == params["id"] {
            Patients = append(Patients[:index], Patients[index+1:]...)
            break
        }
    }
    json.NewEncoder(w).Encode(Patients)
}

// Main function
func main() {
    // Initialize router
    r := mux.NewRouter()

    // Hardcoded data - no database for this example
    Patients = append(Patients, Patient{ID: "1", Firstname: "Andrew", Lastname: "Weber", Address: "22 Hartford, CT", Doctor: &Doctor{Firstname: "John", Lastname: "Doe"}})
    Patients = append(Patients, Patient{ID: "2", Firstname: "Bob", Lastname: "Brown", Address: "26 Hartford, CT", Doctor: &Doctor{Firstname: "Steve", Lastname: "Smith"}})
    fmt.Println(Patients)

    // Route handles & endpoints
    r.Handle("/", isAuthorized(homePage))
    r.HandleFunc("/patients", getPatients).Methods("GET")
    r.HandleFunc("/patients/{id}", getPatient).Methods("GET")
    r.HandleFunc("/patients", addPatient).Methods("POST")
    r.HandleFunc("/patients/{id}", updatePatient).Methods("PUT")
    r.HandleFunc("/patients/{id}", deletePatient).Methods("DELETE")

    // Start server, listen at port 9000
    log.Fatal(http.ListenAndServe(":9000", r))
}
