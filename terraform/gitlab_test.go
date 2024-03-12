package terraform

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	config "github.com/volkovartem/joven/config"
)

func TestMakeGiLabModulesRequest(t *testing.T) {
	server := createMockServer(t)
	defer server.Close()
	config := generateMockConfig()
	// Make the request to the mock server
	responses, _, err := makeGiLabModulesRequest(config, server.URL)

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	expected := &[]Response{

		{
			Name:    "ecs-application/aws",
			Version: "0.0.1",
		},
		{
			Name:    "tgw-module/aws",
			Version: "0.0.1",
		},
	}

	if !reflect.DeepEqual(responses, expected) {
		t.Errorf("Expected %v, got %v", expected, responses)
	}
	fmt.Println(server.URL)

}

func createMockServer(t *testing.T) *httptest.Server {
	t.Helper()
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
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("X-Total-Pages", "3")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, data)
	}))
}

func generateMockConfig() *config.Config{
		// Create a mock config
		return &config.Config{
			Groups: []string{"group-id"},
			Token:  "your-private-token",
		}
}

func TestCreateGitlabUrl(t *testing.T) {
	t.Run("Test First page", func(t *testing.T) {
		config := generateMockConfig()
		url, err := createGitLabUrl(config, "1")
		if err != nil {
			t.Errorf("Unable to generate URL %s", err)
		}
		expected := "https://gitlab.com/api/v4/groups/group-id/packages?package_type=terraform_module&pagination=keyset&page=1&per_page=100&sort=asc"
		if url != expected {
			t.Errorf("got %s want %s given", url, expected)
		}
	})
	t.Run("Test Empty page", func(t *testing.T) {
		config := generateMockConfig()
		_, err := createGitLabUrl(config, "")

		if err != ErrorPageNumberEmpty {
			t.Errorf("got %s want %s given", ErrorPageNumberEmpty, err)
		}
	})
}


