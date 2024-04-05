package cdk

import (
	"os"
	"reflect"
	"testing"
)

func Test_getCDKConstructs(t *testing.T) {
	type args struct {
		fileName string
	}
	tests := []struct {
		name string
		args args
		want []ConstructInfo
	}{
		{
			name: "Test",
			args: args{fileName: "testdata/tree_nostage.json"},
			want: []ConstructInfo{
				{
					Fqn:     "@cdk-platform/dev-deps",
					Version: "0.0.24",
				},
				{
					Fqn:     "aws-cdk-lib",
					Version: "2.113.0",
				},
			},
		},
		{
			name: "Test with stage",
			args: args{fileName: "testdata/tree_stage.json"},
			want: []ConstructInfo{

				{
					Fqn:     "aws-cdk-lib",
					Version: "2.127.0",
				},
				{
					Fqn:     "@cdk-platform/dev-deps",
					Version: "0.0.24",
				},
				{
					Fqn:     "@cdk-platform/security-requirements",
					Version: "0.0.3",
				},
				{
					Fqn:     "constructs",
					Version: "10.3.0",
				},
				{
					Fqn:     "@cdk-platform/finops-practices",
					Version: "0.0.6",
				},
			},
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			file, err := os.ReadFile(tt.args.fileName)
			if err != nil {
				t.Errorf("Error reading file %s: %v", tt.args.fileName, err)
			}
			if got := getCDKConstructs(file); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getNodeCDKConstructs() = %v, want %v", got, tt.want)
			}
		})
	}
}
