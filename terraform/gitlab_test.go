package terraform

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	config "github.com/volkovartem/joven/config"
)

func TestMakeGiLabModulesRequest(t *testing.T) {
	server := createMockServer(t)
	defer server.Close()

	// Create a mock config
	config := &config.Config{
		Groups: []string{"group-name"},
		Token:  "your-private-token",
	}
	// Make the request to the mock server
	resp, err := makeGiLabModulesRequest(config, server.URL)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	fmt.Println(server.URL)

	// Verify the response status code
	expectedStatusCode := http.StatusOK
	if resp.StatusCode != expectedStatusCode {
		t.Errorf("Expected status code %d, got %d", expectedStatusCode, resp.StatusCode)
	}
}

func createMockServer(t *testing.T) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Verify the request method and URL
		if r.Method != "GET" {
			t.Errorf("Expected GET request, got %s", r.Method)
		}
		// Verify the request headers
		expectedToken := "your-private-token"
		if r.Header.Get("PRIVATE-TOKEN") != expectedToken {
			t.Errorf("Expected PRIVATE-TOKEN header %s, got %s", expectedToken, r.Header.Get("PRIVATE-TOKEN"))
		}
		
		data := `[
			{
				"name": "ecs-application/aws",
				"version": "0.0.1"
			},
			{
				"name": "tgw-module/aws",
				"version": "0.0.1"
			}
		
		]`
		err := json.NewEncoder(w).Encode(data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)

	}))
}