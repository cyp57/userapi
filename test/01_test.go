package test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestHelloHandler(t *testing.T) {
	// Create a test HTTP request to the /hello endpoint
	req, err := http.NewRequest("GET", "/HelloHandler", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a recorder to capture the response
	rr := httptest.NewRecorder()

	// Create a Gin router and add the HelloHandler to it
	r := gin.Default()
	// r.GET("/HelloHandler", new(v1.Handler1Receiver).HelloHandler)

	// Perform the request
	r.ServeHTTP(rr, req)

	// Check the response status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Check the response body
	expected := `{"message":"Hello, World!"}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}
