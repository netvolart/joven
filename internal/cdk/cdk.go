package cdk

import (
	"errors"
	"log"
	"os"
)

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
		constructs := getNodeCDKConstructs(localCdkTree)
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

