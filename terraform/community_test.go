package terraform

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/netvolart/joven/iac"
)

func Test_getModuleVersionFromRegistry(t *testing.T) {
	server := createRegistryMockServer(t)
	defer server.Close()

	localModule := LocalModule{
		Source:  "registry.terraform.io/terraform-aws-modules/vpc/aws",
		Version: "5.5.0",
		Type:    "community",
	}

	modules, err := getModuleVersionsFromRegistry(server.URL, localModule)

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	expected := iac.Package{

		Name:          "registry.terraform.io/terraform-aws-modules/vpc/aws",
		LocalVersion:  "5.5.0",
		LatestVersion: "5.6.0",
		Link:          "https://github.com/terraform-aws-modules/terraform-aws-vpc",
	}

	if !reflect.DeepEqual(modules, expected) {
		t.Errorf("Expected %v, got %v", expected, modules)
	}

}

func createRegistryMockServer(t *testing.T) *httptest.Server {
	t.Helper()
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("Expected GET request, got %s", r.Method)
		}

		data := `{
			"id": "terraform-aws-modules/vpc/aws/5.6.0",
			"owner": "antonbabenko",
			"namespace": "terraform-aws-modules",
			"name": "vpc",
			"version": "5.6.0",
			"provider": "aws",
			"provider_logo_url": "/images/providers/aws.png",
			"description": "Terraform module to create AWS VPC resources 🇺🇦",
			"source": "https://github.com/terraform-aws-modules/terraform-aws-vpc",
			"tag": "v5.6.0",
			"published_at": "2024-03-14T13:14:30.519819Z",
			"downloads": 69908067,
			"verified": false,
			"providers": [
			  "aws"
			],
			"versions": [
			  "1.0.0",
			  "5.5.1",
			  "5.5.2",
			  "5.5.3",
			  "5.6.0"
			]
		  }
		  `
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, data)
	}))
}

func TestCreateModuleCommunityUrl(t *testing.T) {
	t.Run("Test First page", func(t *testing.T) {
		url, err := createModuleCommunityUrl("registry.terraform.io/terraform-aws-modules/vpc/aws")
		if err != nil {
			t.Errorf("Unable to generate URL %s", err)
		}
		expected := "https://registry.terraform.io/v1/modules/terraform-aws-modules/vpc/aws"
		if url != expected {
			t.Errorf("got %s want %s given", url, expected)
		}
	})
	t.Run("Test Empty page", func(t *testing.T) {
		_, err := createModuleCommunityUrl("")

		if err != ErrorPageNumberEmpty {
			t.Errorf("got %s want %s given", ErrorPageNumberEmpty, err)
		}
	})
}
