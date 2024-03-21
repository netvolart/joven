package terraform

// import (
// 	"reflect"
// 	"testing"
// )

// func TestGetLocalModules(t *testing.T) {
// 	data := `{
// 		"Modules": [
// 			{
// 				"Key": "",
// 				"Source": "",
// 				"Dir": "."
// 			},
// 			{
// 				"Key": "vpc",
// 				"Source": "gitlab.com/mycorp/vpc-module/aws",
// 				"Version": "0.0.18",
// 				"Dir": ".terraform/modules/vpc"
// 			},
// 			{
// 				"Key": "ecs",
// 				"Source": "gitlab.com/mycorp/ecs-module/aws",
// 				"Version": "0.4.1",
// 				"Dir": ".terraform/modules/ecs"
// 			},
// 			{
// 				"Key": "vpc.vpc",
// 				"Source": "registry.terraform.io/terraform-aws-modules/vpc/aws",
// 				"Version": "5.1.2",
// 				"Dir": ".terraform/modules/vpc.vpc"
// 			}
// 		]
// 	}`

// 	localModules, err := GetLocalModules([]byte(data))
// 	if err != nil {
// 		t.Errorf(err.Error())
// 	}

// 	expected := &LocalModules{
// 		Modules: []LocalModule{
// 			{
// 				Source:  "gitlab.com/mycorp/vpc-module/aws",
// 				Version: "0.0.18",
// 			},
// 			{
// 				Source:  "gitlab.com/mycorp/ecs-module/aws",
// 				Version: "0.4.1",
// 			},
// 			{
// 				Source:  "registry.terraform.io/terraform-aws-modules/vpc/aws",
// 				Version: "5.1.2",
// 			},
// 		},
// 	}
// 	if !reflect.DeepEqual(localModules, expected) {
// 		t.Errorf("Expected %v, got %v", expected, localModules)
// 	}
// }

// func Test_setModuleType(t *testing.T) {

// 	localModules := &LocalModules{
// 		Modules: []LocalModule{
// 			{
// 				Source:  "gitlab.com/mycorp/vpc-module/aws",
// 				Version: "0.0.18",
// 			},
// 			{
// 				Source:  "gitlab.com/mycorp/ecs-module/aws",
// 				Version: "0.4.1",
// 			},
// 			{
// 				Source:  "registry.terraform.io/terraform-aws-modules/vpc/aws",
// 				Version: "5.1.2",
// 			},
// 		},
// 	}

// 	result := setModulesSourceType(localModules)

// 	expected := &LocalModules{
// 		Modules: []LocalModule{
// 			{
// 				Source:  "gitlab.com/mycorp/vpc-module/aws",
// 				Version: "0.0.18",
// 				Type:    "gitlab",
// 			},
// 			{
// 				Source:  "gitlab.com/mycorp/ecs-module/aws",
// 				Version: "0.4.1",
// 				Type:    "gitlab",
// 			},
// 			{
// 				Source:  "registry.terraform.io/terraform-aws-modules/vpc/aws",
// 				Version: "5.1.2",
// 				Type:    "community",
// 			},
// 		},
// 	}

// 	if !reflect.DeepEqual(result, expected) {
// 		t.Errorf("Expected %v, got %v", expected, result)
// 	}

// }
