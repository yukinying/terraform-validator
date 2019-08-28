package tfgcv

import (
	"path/filepath"
	"testing"

	"github.com/GoogleCloudPlatform/terraform-validator/tfplan"
)

const testDataDir = "../test/read_planned_assets"

func TestReadPlannedAssets(t *testing.T) {
	const testProjectName = "gl-akopachevskyy-sql-db"
	const testAncestryName = "ancestry"
	type args struct {
		file      string
		project   string
		ancestry  string
		tfVersion string
		offline   bool
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			"Test TF12 and JSON plan",
			args{"tf12plan.json", testProjectName, testAncestryName, tfplan.TF12, true},
			2,
			false,
		},
		{
			"Test TF12 and binary plan",
			args{"tf12plan.json", testProjectName, testAncestryName, tfplan.TF11, true},
			0,
			true,
		},
		{
			"Test TF11 and JSON plan",
			args{"tf11plan.tfplan", testProjectName, testAncestryName, tfplan.TF12, true},
			0,
			true,
		},
		// Note: Terraform tfplan parsing is not supported in v0.12.
		// The following test is commented out as the package imported is also in v0.12 and
		// the test will always fail.
		//
		// To test the logic, please checkout the branch terraform-v0.11
		// {
		// 	"Test TF11 and binary plan",
		// 	args{"tf11plan.tfplan", testProjectName, testAncestryName, tfplan.TF11, true},
		// 	2,
		// 	false,
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testFile := filepath.Join(testDataDir, tt.args.file)
			got, err := ReadPlannedAssets(testFile, tt.args.project, tt.args.ancestry, tt.args.tfVersion, tt.args.offline)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadPlannedAssets() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(got) != tt.want {
				t.Errorf("ReadPlannedAssets() = %v, want %v", len(got), tt.want)
			}
		})
	}
}
