package cdk

import (
	"fmt"
	"log"
	"os"
	"reflect"
	"testing"
)

func createTestFile(t *testing.T, dir, name, content string) (filePath string) {
	t.Helper()
	err := os.MkdirAll(dir, 0755)
	if err != nil {
		log.Fatal(err)
	}

	err = os.WriteFile(name, []byte(content), 0644)
	if err != nil {
		t.Fatal(err)
	}
	return fmt.Sprintf("%s/%s", dir, name)

}

func Test_detectLanguage(t *testing.T) {
	tests := []struct {
		name     string
		dir      string
		fileType string
		language string
		content  string
	}{
		{
			name:     "TypeScript test",
			dir:      ".",
			fileType: "package.json",
			language: "typescript",
			content: `
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
			`,
		},
		{
			name:     ".NET test",
			dir:      "infra",
			fileType: "infra/project.csproj",
			language: "dotnet",
			content: `<Project Sdk="Microsoft.NET.Sdk">

			<PropertyGroup>
			  <OutputType>Exe</OutputType>
			  <TargetFramework>net6.0</TargetFramework>
			  <!-- Roll forward to future major versions of the netcoreapp as needed -->
			  <RollForward>Major</RollForward>
			  <ValidateExecutableReferencesMatchSelfContained>false</ValidateExecutableReferencesMatchSelfContained>
			  <RootNamespace>TestApp.Cdk</RootNamespace>
			  <NoWarn>CA1806</NoWarn>
			</PropertyGroup>
		  
			<ItemGroup>
			  <PackageReference Include="Constructs" Version="10.3.0" />
			  <PackageReference Include="Company.Cdk.Lib.CdkCustomAuthorizer" Version="0.1.34" />
			  <PackageReference Include="Company.Cdk.Otel" Version="0.1.34" />
			  <PackageReference Include="Company.Cdk.Shared" Version="0.1.34" />
			  <PackageReference Include="Amazon.Jsii.Analyzers" Version="1.71.0" PrivateAssets="all" />
			  <PackageReference Include="AWSSDK.SSO" Version="3.7.200.14" />
			  <PackageReference Include="AWSSDK.SSOOIDC" Version="3.7.200.14" />
			</ItemGroup>
		  
			<ItemGroup>
			  <ProjectReference Include="..\..\src\NpsApp.Domain\TestApp.Domain.csproj" />
			</ItemGroup>
		  </Project>
		  `,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			createTestFile(t, tt.dir, tt.fileType, tt.content)

			got, err := detectLanguage()
			if err != nil {
				t.Errorf("detectLanguage error = %v", err)
				return
			}
			if !reflect.DeepEqual(got, tt.language) {
				t.Errorf("detectLanguage() = %v, want %v", got, tt.language)
			}
			os.Remove(tt.fileType)

			if tt.dir != "." {
				os.RemoveAll(tt.dir)
			}
		})
	}
}
