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

var destroyCmd = &cobra.Command{
	Use:   "destroy",
	Short: "Destroy provisioned infrastructure",
	Long: `Destroys all infrastructure resources by running 'terraform destroy'.

⚠️  WARNING: This is a destructive operation that cannot be undone!

Requires:
  - Terraform binary installed locally
  - Existing Terraform state (resources must be provisioned first)
  - Cloud credentials configured`,
	RunE: runDestroy,
}

func runDestroy(cmd *cobra.Command, args []string) error {
	// Check if infra directory exists
	if _, err := os.Stat("infra"); os.IsNotExist(err) {
		return fmt.Errorf("infra/ directory not found")
	}

	// Confirm destruction
	fmt.Println("⚠️  WARNING: This will DESTROY all provisioned infrastructure!")
	fmt.Print("Type 'destroy' to confirm: ")

	reader := bufio.NewReader(os.Stdin)
	response, err := reader.ReadString('\n')
	if err != nil {
		return fmt.Errorf("failed to read input: %w", err)
	}

	response = strings.TrimSpace(strings.ToLower(response))
	if response != "destroy" {
		fmt.Println("Aborted.")
		return nil
	}

	// Run terraform destroy
	tfDestroy := exec.Command("terraform", "destroy")
	tfDestroy.Dir = "infra"
	tfDestroy.Stdout = os.Stdout
	tfDestroy.Stderr = os.Stderr
	tfDestroy.Stdin = os.Stdin

	fmt.Println("\nDestroying infrastructure...")
	if err := tfDestroy.Run(); err != nil {
		return fmt.Errorf("terraform destroy failed: %w", err)
	}

	fmt.Println("\n✓ Infrastructure destroyed successfully")

	return nil
}
