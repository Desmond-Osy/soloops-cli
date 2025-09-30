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

	"github.com/soloops/soloops-cli/pkg/config"
	"github.com/spf13/cobra"
)

var validateCmd = &cobra.Command{
	Use:   "validate",
	Short: "Validate soloops.yaml configuration",
	Long: `Parses and validates the soloops.yaml manifest.

Checks for:
  - Valid YAML syntax
  - Required fields (project, cloud, environments)
  - Budget constraints
  - Blueprint configurations
  - Policy settings

Returns detailed error messages with suggestions for fixing issues.`,
	RunE: runValidate,
}

func runValidate(cmd *cobra.Command, args []string) error {
	cfg, err := config.Load(configFile)
	if err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}

	if err := cfg.Validate(); err != nil {
		return fmt.Errorf("validation failed: %w", err)
	}

	fmt.Printf("✓ Configuration is valid (%s)\n", configFile)
	fmt.Printf("  Project: %s\n", cfg.Project)
	fmt.Printf("  Cloud: %s\n", cfg.Cloud)
	fmt.Printf("  Environments: %d\n", len(cfg.Environments))
	for _, env := range cfg.Environments {
		fmt.Printf("    - %s (%s): $%.2f budget, %d blueprints\n",
			env.Name, env.Region, env.BudgetUSD, len(env.Blueprints))
	}

	return nil
}