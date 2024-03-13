package terraform

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	config "github.com/volkovartem/joven/config"
)

// Create a mock config
func generateMockConfig(t *testing.T) *config.Config {
	t.Helper()
	return &config.Config{
		Groups: []string{"group-id"},
		Token:  "your-private-token",
	}
}

func TestCreateModuleGitlabUrl(t *testing.T) {
	t.Run("Test First page", func(t *testing.T) {
		config := generateMockConfig(t)
		url, err := createModuleGitlabUrl(config, "ecs-module/aws")
		if err != nil {
			t.Errorf("Unable to generate URL %s", err)
		}
		expected := "https://gitlab.com/api/v4/groups/group-id/packages?package_type=terraform_module&package_name=ecs-module/aws&sort=asc"
		if url != expected {
			t.Errorf("got %s want %s given", url, expected)
		}
	})
	t.Run("Test Empty page", func(t *testing.T) {
		config := generateMockConfig(t)
		_, err := createModuleGitlabUrl(config, "")

		if err != ErrorPageNumberEmpty {
			t.Errorf("got %s want %s given", ErrorPageNumberEmpty, err)
		}
	})
}

func Test_getModuleVersionsFromGitLab(t *testing.T) {
	server := createMockServer(t)
	defer server.Close()
	config := generateMockConfig(t)
	// Make the request to the mock server
	modules, err := getModuleVersionsFromGitLab(config, server.URL)

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	expected := []*TerraformModule{
		{
			Name:          "ecs-module/aws",
			LocalVersion:  "",
			LatestVersion: "0.0.1",
			Link:          "https://gitlab.com/mygroup/terraformmodules/ModuleECS/-/infrastructure_registry/234245",
		},
		{
			Name:          "ecs-module/aws",
			LocalVersion:  "",
			LatestVersion: "0.3.0",
			Link:          "https://gitlab.com/mygroup/terraformmodules/ModuleECS/-/infrastructure_registry/353555",
		},
	}

	if !reflect.DeepEqual(modules, expected) {
		t.Errorf("Expected %v, got %v", expected, modules)
	}

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
				"name": "ecs-module/aws",
				"version": "0.0.1",
				"_links": {
					"web_path": "/mygroup/terraformmodules/ModuleECS/-/infrastructure_registry/234245",
					"delete_api_path": "https://gitlab.com/api/v4/projects/3423266/packages/234245"
				  }
			},
			{
				"name": "ecs-module/aws",
				"version": "0.3.0",
				"_links": {
					"web_path": "/mygroup/terraformmodules/ModuleECS/-/infrastructure_registry/353555",
					"delete_api_path": "https://gitlab.com/api/v4/projects/3423266/packages/4353553"
				  }
			}
		
		]`
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, data)
	}))
}

func Test_clearOldVersions(t *testing.T) {
	tests := []struct {
		name    string
		modules []*TerraformModule
		want    []*TerraformModule
		wantErr bool
	}{
		{
			name: "Normal test",
			modules: []*TerraformModule{
				{
					Name:          "ecs-application/aws",
					LatestVersion: "0.0.1",
				},
				{
					Name:          "tgw-module/aws",
					LatestVersion: "0.0.2",
				},
				{
					Name:          "tgw-module/aws",
					LatestVersion: "0.3.1",
				},
				{
					Name:          "tgw-module/aws",
					LatestVersion: "1.0.1",
				},
			},
			want: []*TerraformModule{
				{
					Name:          "ecs-application/aws",
					LatestVersion: "0.0.1",
				},
				{
					Name:          "tgw-module/aws",
					LatestVersion: "1.0.1",
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := clearOldVersions(tt.modules)
			if (err != nil) != tt.wantErr {
				t.Errorf("clearOldVersions() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("clearOldVersions() = %v, want %v", got, tt.want)
			}
		})
	}
}
