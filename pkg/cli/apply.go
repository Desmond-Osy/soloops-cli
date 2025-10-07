// Copyright 2025 SoloOps Contributors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cli

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

var autoApprove bool

var applyCmd = &cobra.Command{
	Use:   "apply",
	Short: "Apply infrastructure changes",
	Long: `Provisions infrastructure by running 'terraform apply'.

Requires:
  - Terraform binary installed locally
  - Generated Terraform files (run 'soloops generate' first)
  - Cloud credentials configured

Flags:
  --auto-approve: Skip interactive approval prompt`,
	RunE: runApply,
}

func init() {
	applyCmd.Flags().BoolVar(&autoApprove, "auto-approve", false, "Skip interactive approval prompt")
}

func runApply(cmd *cobra.Command, args []string) error {
	// Check if infra directory exists
	if _, err := os.Stat("infra"); os.IsNotExist(err) {
		return fmt.Errorf("infra/ directory not found. Run 'soloops generate' first")
	}

	// Initialize Terraform if needed
	tfInit := exec.Command("terraform", "init")
	tfInit.Dir = "infra"
	tfInit.Stdout = os.Stdout
	tfInit.Stderr = os.Stderr

	fmt.Println("Initializing Terraform...")
	if err := tfInit.Run(); err != nil {
		return fmt.Errorf("terraform init failed: %w", err)
	}

	// Confirm before applying
	if !autoApprove {
		fmt.Println("\n⚠️  This will provision real infrastructure and may incur costs.")
		fmt.Print("Do you want to continue? (yes/no): ")

		reader := bufio.NewReader(os.Stdin)
		response, err := reader.ReadString('\n')
		if err != nil {
			return fmt.Errorf("failed to read input: %w", err)
		}

		response = strings.TrimSpace(strings.ToLower(response))
		if response != "yes" && response != "y" {
			fmt.Println("Aborted.")
			return nil
		}
	}

	// Run terraform apply
	applyArgs := []string{"apply"}
	if autoApprove {
		applyArgs = append(applyArgs, "-auto-approve")
	}

	tfApply := exec.Command("terraform", applyArgs...)
	tfApply.Dir = "infra"
	tfApply.Stdout = os.Stdout
	tfApply.Stderr = os.Stderr
	tfApply.Stdin = os.Stdin

	fmt.Println("\nApplying infrastructure changes...")
	if err := tfApply.Run(); err != nil {
		return fmt.Errorf("terraform apply failed: %w", err)
	}

	fmt.Println("\n✓ Infrastructure provisioned successfully")
	fmt.Println("\nTo view outputs, run: cd infra && terraform output")
	fmt.Println("To destroy resources, run: soloops destroy")

	return nil
}
