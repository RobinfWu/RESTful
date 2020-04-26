package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
    "time"
    jwt "github.com/dgrijalva/jwt-go"
)

// In practice, use environmental variables
// Ex: var mySigningKey = os.Get("MY_JWT_TOKEN")
var mySigningKey = []byte("secretkey")

func homePage(w http.ResponseWriter, r *http.Request) {
    validToken, err := GenerateJWT()
    if err != nil {
        fmt.Println("Failed to generate token")
    }
    //fmt.Println(w, validToken)
    client := &http.Client{}
    req, _ := http.NewRequest("GET", "http://localhost:9000/", nil)
    req.Header.Set("Token", validToken)
    res, err := client.Do(req)
    if err != nil {
        fmt.Fprintf(w, "Error: %s", err.Error())
    }

    body, err := ioutil.ReadAll(res.Body)
    if err != nil {
        fmt.Println(err)
    }
    fmt.Fprintf(w, string(body))
}

func GenerateJWT() (string, error) {
    token := jwt.New(jwt.SigningMethodHS256)

    claims := token.Claims.(jwt.MapClaims)

    claims["authorized"] = true
    claims["client"] = "John Doe"
    claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

    tokenString, err := token.SignedString(mySigningKey)

    if err != nil {
        fmt.Errorf("Something Went Wrong: %s", err.Error())
        return "", err
    }

    return tokenString, nil
}

func handleRequests() {
    http.HandleFunc("/", homePage)
    log.Fatal(http.ListenAndServe(":8001", nil))
}

func main() {
    fmt.Println("Simple Client")
    handleRequests()
}