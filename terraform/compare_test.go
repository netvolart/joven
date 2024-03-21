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
