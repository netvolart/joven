package cdk

import (
	"reflect"
	"sort"
	"testing"
)

func sortPackages(t *testing.T, packages []CDKPackage) {
	t.Helper()
	sort.Slice(packages, func(i, j int) bool {
		return packages[i].Name < packages[j].Name
	})
}

func Test_parsePackageJson(t *testing.T) {
	data := `
	{
		"name": "@test/core-infrastructure",
		"version": "0.1.1",
		"bin": {
		  "core-infrastructure": "bin/core-infrastructure.js"
		},
		"scripts": {
		  "build": "tsc",
		  "watch": "tsc -w",
		  "test": "jest",
		  "cdk": "cdk",
		  "cdk-diff-dev": "cdk diff -c config=dev1",
		  "cdk-synth-dev": "cdk synth -c config=dev1",
		  "cdk-deploy-dev": "cdk deploy -c config=dev1"
		},
		"devDependencies": {
		  "@semantic-release/changelog": "^6.0.3",
		  "@semantic-release/commit-analyzer": "^11.1.0",
		  "@semantic-release/gitlab": "^13.0.2",
		  "@semantic-release/release-notes-generator": "^12.1.0",
		  "@types/jest": "^29.5.3",
		  "@types/node": "20.4.9",
		  "aws-cdk": "^2.127.0",
		  "jest": "^29.6.2",
		  "semantic-release": "^23.0.2",
		  "semver": "^7.6.0",
		  "ts-jest": "^29.1.1",
		  "ts-node": "^10.9.1",
		  "typedoc": "^0.25.1",
		  "typedoc-plugin-markdown": "^3.16.0",
		  "typescript": "~5.1.6"
		},
		"dependencies": {
		  "@aws-cdk/aws-msk-alpha": "^2.100.0-alpha.0",
		  "@cdk-platform/network-settings": "^0.0.25",
		  "@aws-sdk/client-ec2": "^3.465.0",
		  "@cdk-platform/dev-deps": "^0.0.24",
		  "aws-cdk-lib": "2.127.0",
		  "cdk-dia": "^0.10.0",
		  "cdk-ecr-deployment": "^2.5.30",
		  "cdk-nag": "^2.27.173",
		  "constructs": "^10.0.0",
		  "source-map-support": "^0.5.21"
		}
	  }	  
	`

	packages := parsePackageJson([]byte(data))

	sortPackages(t, packages)

	expected := []CDKPackage{
		{
			Name:    "@semantic-release/changelog",
			Version: "^6.0.3",
		},
		{
			Name:    "@semantic-release/commit-analyzer",
			Version: "^11.1.0",
		},
		{
			Name:    "@semantic-release/gitlab",
			Version: "^13.0.2",
		},
		{
			Name:    "@semantic-release/release-notes-generator",
			Version: "^12.1.0",
		},
		{
			Name:    "@types/jest",
			Version: "^29.5.3",
		},
		{
			Name:    "@types/node",
			Version: "20.4.9",
		},
		{
			Name:    "aws-cdk",
			Version: "^2.127.0",
		},
		{
			Name:    "jest",
			Version: "^29.6.2",
		},
		{
			Name:    "semantic-release",
			Version: "^23.0.2",
		},
		{
			Name:    "semver",
			Version: "^7.6.0",
		},
		{
			Name:    "ts-jest",
			Version: "^29.1.1",
		},
		{
			Name:    "ts-node",
			Version: "^10.9.1",
		},
		{
			Name:    "typedoc",
			Version: "^0.25.1",
		},
		{
			Name:    "typedoc-plugin-markdown",
			Version: "^3.16.0",
		},
		{
			Name:    "typescript",
			Version: "~5.1.6",
		},
		{
			Name:    "@aws-cdk/aws-msk-alpha",
			Version: "^2.100.0-alpha.0",
		},
		{
			Name:    "@cdk-platform/network-settings",
			Version: "^0.0.25",
		},
		{
			Name:    "@aws-sdk/client-ec2",
			Version: "^3.465.0",
		},
		{
			Name:    "@cdk-platform/dev-deps",
			Version: "^0.0.24",
		},
		{
			Name:    "aws-cdk-lib",
			Version: "2.127.0",
		},
		{
			Name:    "cdk-dia",
			Version: "^0.10.0",
		},
		{
			Name:    "cdk-ecr-deployment",
			Version: "^2.5.30",
		},
		{
			Name:    "cdk-nag",
			Version: "^2.27.173",
		},
		{
			Name:    "constructs",
			Version: "^10.0.0",
		},
		{
			Name:    "source-map-support",
			Version: "^0.5.21",
		},
	}
	sortPackages(t, expected)

	if !reflect.DeepEqual(packages, expected) {
		t.Errorf("Expected %v, got %v", expected, packages)
	}
}

func Test_clearFqn(t *testing.T) {
	fqn := "@cdk-platform/network-settings.WebACLStack"
	expected := "@cdk-platform/network-settings"
	result := clearFqn(fqn)
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}

// func Test_getCDKConstructs(t *testing.T) {

// 	getNodeCDKConstructs()

// 	// if !reflect.DeepEqual(constructs, expected) {
// 	// 	t.Errorf("Expected %v, got %v", expected, packages)
// 	// }

// }

func Test_removeDuplicates(t *testing.T) {
	type args struct {
		packages []CDKPackage
	}
	tests := []struct {
		name string
		args args
		want []CDKPackage
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeDuplicates(tt.args.packages); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeDuplicates() = %v, want %v", got, tt.want)
			}
		})
	}
}
