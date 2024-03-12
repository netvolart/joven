package terraform

import (
	"encoding/json"
	"fmt"
	"os"
)
type LocalModule struct {
    // Key     string `json:"Key"`
    Source  string `json:"Source"`
    Version string `json:"Version"`
    // Dir     string `json:"Dir"`
}

type LocalModules struct {
    Modules []LocalModule `json:"Modules"`
}


func GetLocalModules() (*LocalModules, error) {
	configData, err := os.ReadFile(".terraform/modules/modules.json")
	if err != nil {
		return nil, err
	}
	var localModules LocalModules
	err = json.Unmarshal(configData, &localModules)
	if err != nil {
		return nil, err
	}
	for _, module := range localModules.Modules {
		fmt.Printf("%v: %v\n", module.Source, module.Version)
	}
	return &localModules, nil
}