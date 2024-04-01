package cdk

import (
	"encoding/json"
	"errors"
	"log"
	"os"
	"strings"
)

type TreeNode struct {
	ID            string              `json:"id"`
	Path          string              `json:"path"`
	ConstructInfo ConstructInfo       `json:"constructInfo"`
	Child         map[string]TreeNode `json:"children,omitempty"`
}

type Tree struct {
	Version string   `json:"version"`
	Tree    TreeNode `json:"tree"`
}

type ConstructInfo struct {
	Fqn     string `json:"fqn"`
	Version string `json:"version"`
}

func unmarshalTree(data []byte) Tree {
	tree := Tree{}
	err := json.Unmarshal(data, &tree)
	if err != nil {
		log.Fatalf("error unmarshaling tree JSON: %v", err)
	}
	return tree
}

func clearFqn(Fqn string) string {
	lib := strings.Split(Fqn, ".")[0]
	if lib != "" {
		return lib
	}
	return ""
}

func removeDuplicatesConstructInfo(constructs []ConstructInfo) []ConstructInfo {
	seen := make(map[ConstructInfo]bool)
	uniqueConstructs := []ConstructInfo{}
	for _, entry := range constructs {
		construct := ConstructInfo{
			Fqn:     clearFqn(entry.Fqn),
			Version: entry.Version,
		}
		if !seen[construct] {
			seen[construct] = true
			uniqueConstructs = append(uniqueConstructs, construct)
		}
	}
	return uniqueConstructs
}

func findConstructInfo(childs map[string]TreeNode) (constructs []ConstructInfo) {

	for _, child := range childs {
		if child.ID == "Default" {
			return constructs
		}
		constructs = append(constructs, child.ConstructInfo)

		constr := findConstructInfo(child.Child)
		constructs = append(constructs, constr...)

	}

	return constructs
}

func getCDKConstructs(data []byte) (uniqueConstructs []ConstructInfo) {
	// Parse tree JSON
	tree := unmarshalTree(data)
	// Find all constructs in the tree recursively
	constructs := findConstructInfo(tree.Tree.Child)

	uniqueConstructs = append(uniqueConstructs, constructs...)
	uniqueConstructs = removeDuplicatesConstructInfo(uniqueConstructs)

	return
}

func CompareCDKConstructs() ([]CDKPackage, error) {
	lang, err := detectLanguage()
	if err != nil {
		log.Fatalf("error detecting language: %v", err)
	}
	switch lang {
	case "typescript":
		localCdkTree, err := os.ReadFile("cdk.out/tree.json")
		if err != nil {
			return nil, err
		}
		constructs := getCDKConstructs(localCdkTree)
		packages := formNodeCDKPackages(constructs)
		packagesWithoutDuplicates := removeDuplicates(packages)
		for k, pack := range packagesWithoutDuplicates {
			pack.getPackageNpmInfo()
			pack.setNpmPackageType()
			pack.setOutdated()
			packagesWithoutDuplicates[k] = pack
		}
		return packagesWithoutDuplicates, nil
	case "dotnet":
		log.Fatal("dotnet is not supported yet")
	default:
		return nil, errors.New("language is not supported")
	}
	return nil, errors.New("language is not supported")

}
