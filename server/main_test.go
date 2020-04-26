package main

import (
	"testing"
	"bytes"
	"net/http"
	"net/http/httptest"
)

// Check status of GET request with the getPatients()
func TestGetPatients(t *testing.T) {
	request, err := http.NewRequest("GET", "/patients", nil)
	if err != nil {
		t.Fatal(err)
	}

	response := httptest.NewRecorder()
	handler := http.HandlerFunc(getPatients)
	handler.ServeHTTP(response, request)
	if status := response.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

// Check status of GET request with the getPatient() with patient ID 1
func TestGetPatient(t *testing.T) {
	request, err := http.NewRequest("GET", "/patients", nil)

	if err != nil {
		t.Fatal(err)
	}

	q := request.URL.Query()
	q.Add("id", "1")
	request.URL.RawQuery = q.Encode()

	response := httptest.NewRecorder()
	handler := http.HandlerFunc(getPatient)
	handler.ServeHTTP(response, request)

	if status := response.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

// Check status of POST request with the addPatient()
func TestAddPatient(t *testing.T) {
	var jsonStr = []byte(`{"id":"3","firstname":"Alice","lastname":"Endyke","address":"54 Hartford, CT","doctor":{"firstname":"John","lastname":"Stewart"}}`)

	request, err := http.NewRequest("POST", "/patients", bytes.NewBuffer(jsonStr))

	if err != nil {
		t.Fatal(err)
	}

	request.Header.Set("Content-Type", "application/json")
	response := httptest.NewRecorder()
	handler := http.HandlerFunc(addPatient)
	handler.ServeHTTP(response, request)

	if status := response.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

// Check status of PUT request with the updatePatient()
func TestUpdatePatient(t *testing.T) {
	var jsonStr = []byte(`{"id":"3","firstname":"Alice","lastname":"Endyke","address":"43 Treeland, CA","doctor":{"firstname":"John","lastname":"Stewart"}}`)

	request, err := http.NewRequest("PUT", "/patients", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	request.Header.Set("Content-Type", "application/json")
	response := httptest.NewRecorder()
	handler := http.HandlerFunc(updatePatient)
	handler.ServeHTTP(response, request)
	if status := response.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

// Check status of DELETE request with the deletePatient()
func TestDeletePatient(t *testing.T) {
	request, err := http.NewRequest("DELETE", "/patients", nil)
	if err != nil {
		t.Fatal(err)
	}
	q := request.URL.Query()
	q.Add("id", "1")
	request.URL.RawQuery = q.Encode()
	response := httptest.NewRecorder()
	handler := http.HandlerFunc(deletePatient)
	handler.ServeHTTP(response, request)
	if status := response.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}
