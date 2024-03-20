package terraform

import (
	"reflect"
	"testing"
)

func Test_findOutdated(t *testing.T) {
	input := []*TerraformModule{
		{
			Name:          "ecs-application/aws",
			LocalVersion:  "0.0.2",
			LatestVersion: "3.0.0",
			Outdated:      false,
		},
		{
			Name:          "tgw-module/aws",
			LocalVersion:  "0.0.2",
			LatestVersion: "0.0.2",
			Outdated:      false,
		},
	}
	expectedResult := []*TerraformModule{
		{
			Name:          "ecs-application/aws",
			LocalVersion:  "0.0.2",
			LatestVersion: "3.0.0",
			Outdated:      true,
		},
		{
			Name:          "tgw-module/aws",
			LocalVersion:  "0.0.2",
			LatestVersion: "0.0.2",
			Outdated:      false,
		},
	}

	result, err := findOutdated(input)
	if err != nil {
		t.Errorf(err.Error())
	}

	if !reflect.DeepEqual(result, expectedResult) {
		t.Errorf("Expected %v, got %v", expectedResult, result)
	}
}

// func Test_getModuleVersionsFromGitLab(t *testing.T) {
// 	server := createMockServer(t)
// 	defer server.Close()
// 	config := generateMockConfig(t)
// 	// Make the request to the mock server
// 	modules, err := getModuleVersionsFromGitLab(config, server.URL)

// 	if err != nil {
// 		t.Errorf("Unexpected error: %v", err)
// 	}
// 	expected := []*TerraformModule{
// 		{
// 			Name:          "ecs-module/aws",
// 			LocalVersion:  "",
// 			LatestVersion: "0.0.1",
// 			Link:          "https://gitlab.com/mygroup/terraformmodules/ModuleECS/-/infrastructure_registry/234245",
// 		},
// 		{
// 			Name:          "ecs-module/aws",
// 			LocalVersion:  "",
// 			LatestVersion: "0.3.0",
// 			Link:          "https://gitlab.com/mygroup/terraformmodules/ModuleECS/-/infrastructure_registry/353555",
// 		},
// 	}

// 	if !reflect.DeepEqual(modules, expected) {
// 		t.Errorf("Expected %v, got %v", expected, modules)
// 	}

// }