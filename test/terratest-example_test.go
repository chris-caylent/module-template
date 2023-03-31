// This Terratest file is here to serve as an example of a test structure
// You may use bits and pieces of this to suit your needs
// RUNNING THIS EXAMPLE WILL NOT WORK WITHOUT TERRAFORM CODE IN THE EXAMPLE DIRECTORY

package test

import (
	"fmt"
	"path/filepath"
	"testing"

	// use the below package to leverage AWS API calls within your Terratest
	//"github.com/gruntwork-io/terratest/modules/aws"
	"github.com/gruntwork-io/terratest/modules/terraform"
	test_structure "github.com/gruntwork-io/terratest/modules/test-structure"
	//"github.com/stretchr/testify/assert"
)

// example below is the simplest of tests that takes the code in the examples dir
// init, plans, applies, and applies again to check for idempotency
// test will pass if IaC is idempotent and will fail if not, or any other error during apply.

func TestSimpleTest(t *testing.T) {
		exampleName := "terratest-example" // this string should match the name of the directory `examples/simple`
		
		t.Run(exampleName, func(t *testing.T) {
			t.Parallel()
			examplesPath := fmt.Sprintf("examples/%s", exampleName)
			examplesDir := test_structure.CopyTerraformFolderToTemp(t, "../", examplesPath)
			planFilePath := filepath.Join(examplesDir, fmt.Sprintf("%s.txt", exampleName))
			terraformOptions := &terraform.Options{
				NoColor:      true, //prevents CLI output from having standard color output
				TerraformDir: examplesDir,
			}
			planTerraformOptions := &terraform.Options{
				NoColor:      true, //prevents CLI output from having standard color outpu
				PlanFilePath: planFilePath,
				TerraformDir: examplesDir,
			}

			// Generate the plan, show it, and store it for any validation; will error & fail the test if an error is thrown at any stage
			_, err := terraform.InitAndPlanAndShowWithStructE(t, planTerraformOptions)
			if err != nil {
				t.Error(err)
				t.FailNow()
			}

			// don't destroy until there is an error, or the end of the test.
			defer terraform.Destroy(t, terraformOptions)
			
			// Init, plan, apply again, will error if changes to the configuration or detected
			_, err = terraform.InitAndApplyAndIdempotentE(t, terraformOptions)
			if err != nil {
				terraform.Destroy(t, terraformOptions)
				t.FailNow()
			}
		})
	}

