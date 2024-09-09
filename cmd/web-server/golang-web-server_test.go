package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"web_server/configs"
)

func helloString(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hi")
}
func TestWebServerConnection(t *testing.T) {
	// Create a new request to the server
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a response recorder to record the response
	rr := httptest.NewRecorder()

	// Call the handler function or file server
	if useHandlerFunctions {
		handler := http.HandlerFunc(helloString)
		handler.ServeHTTP(rr, req)
	} else {
		http.FileServer(http.Dir(configs.SERVER_STATIC_DIRECTORY)).ServeHTTP(rr, req)
	}

	// Check the response status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Check the response body
	expected := "hi"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}
