package terraform

import (
	"reflect"
	"testing"
)

func TestMergeModules(t *testing.T) {
	gitlabModules := []*TerraformModule{
		{
			Name:          "ecs-application/aws",
			LatestVersion: "0.0.1",
		},
		{
			Name:          "tgw-module/aws",
			LatestVersion: "0.0.1",
		},
	}

	localModules := LocalModules{
		Modules: []LocalModule{
			{
				Source:  "gitlab.com/group/repo/ecs-application/aws",
				Version: "0.0.2",
			},
			{
				Source:  "",
				Version: "0.0.1",
			},
			{
				Source:  "github.com/user/repo/tgw-module/aws",
				Version: "0.0.1",
			},
			{
				Source:  "github.com/user/repo/other-module",
				Version: "1.0.0",
			},
		},
	}

	expectedResult := []*TerraformModule{
		{
			Name:          "ecs-application/aws",
			LocalVersion:  "0.0.2",
			LatestVersion: "0.0.1",
		},
	}

	result := MergeModules(gitlabModules, localModules)

	if !reflect.DeepEqual(result, expectedResult) {
		t.Errorf("Expected %v, got %v", expectedResult, result)
	}
}

func Test_returnOutdated(t *testing.T) {
	input := []*TerraformModule{
		{
			Name:          "ecs-application/aws",
			LocalVersion:  "0.0.2",
			LatestVersion: "3.0.0",
		},
		{
			Name:          "tgw-module/aws",
			LocalVersion:  "0.0.2",
			LatestVersion: "0.0.2",
		},
	}
	expectedResult := []*TerraformModule{
		{
			Name:          "ecs-application/aws",
			LocalVersion:  "0.0.2",
			LatestVersion: "3.0.0",
		},
	}

	result, err := returnOutdated(input)
	if err != nil {
		t.Errorf(err.Error())
	}

	if !reflect.DeepEqual(result, expectedResult) {
		t.Errorf("Expected %v, got %v", expectedResult, result)
	}
}