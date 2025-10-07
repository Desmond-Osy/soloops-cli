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
	"fmt"
	"os"

	"github.com/soloops/soloops-cli/pkg/config"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a new SoloOps project",
	Long: `Creates a starter soloops.yaml manifest with sensible defaults.

The generated manifest includes:
  - Project configuration
  - One production environment
  - Budget settings
  - Sample blueprints (web API, static site)
  - Security policies`,
	RunE: runInit,
}

func runInit(cmd *cobra.Command, args []string) error {
	// Check if file already exists
	if _, err := os.Stat(configFile); err == nil {
		return fmt.Errorf("file already exists: %s (use --file to specify a different path)", configFile)
	}

	// Write the default template
	template := config.DefaultTemplate()
	if err := os.WriteFile(configFile, []byte(template), 0644); err != nil {
		return fmt.Errorf("failed to write config file: %w", err)
	}

	fmt.Printf("✓ Created %s\n", configFile)
	fmt.Println("\nNext steps:")
	fmt.Println("  1. Edit soloops.yaml to customize your infrastructure")
	fmt.Println("  2. Run 'soloops validate' to check your configuration")
	fmt.Println("  3. Run 'soloops generate' to create Terraform files")
	fmt.Println("  4. Run 'soloops preview' to see planned changes")
	fmt.Println("  5. Run 'soloops apply' to provision infrastructure")

	return nil
}
