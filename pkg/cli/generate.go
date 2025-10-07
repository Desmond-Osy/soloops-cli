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

	"github.com/OplexTech/soloops-cli/pkg/config"
	"github.com/OplexTech/soloops-cli/pkg/generator"
	"github.com/spf13/cobra"
)

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate Terraform infrastructure code",
	Long: `Reads soloops.yaml and generates Terraform files in the infra/ directory.

Generates:
  - provider.tf - Cloud provider configuration
  - main.tf - Main infrastructure resources
  - variables.tf - Input variables
  - outputs.tf - Output values
  - budget.tf - Budget alerts

Supports blueprints:
  - web_api: Serverless API (Lambda, API Gateway, WAF)
  - static_site: Static website (S3, CloudFront, HTTPS)
  - database: Managed databases (RDS, Aurora Serverless)`,
	RunE: runGenerate,
}

func runGenerate(cmd *cobra.Command, args []string) error {
	cfg, err := config.Load(configFile)
	if err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}

	if err := cfg.Validate(); err != nil {
		return fmt.Errorf("validation failed: %w", err)
	}

	// Determine target environment
	targetEnv := envName
	if targetEnv == "" {
		targetEnv = cfg.Environments[0].Name
		fmt.Printf("Using default environment: %s\n", targetEnv)
	}

	env, err := cfg.GetEnvironment(targetEnv)
	if err != nil {
		return err
	}

	// Generate Terraform code
	gen := generator.New(cfg, env)
	if err := gen.Generate(); err != nil {
		return fmt.Errorf("generation failed: %w", err)
	}

	fmt.Printf("✓ Generated Terraform files in infra/\n")
	fmt.Printf("  Environment: %s (%s)\n", env.Name, env.Region)
	fmt.Printf("  Budget: $%.2f/month\n", env.BudgetUSD)
	fmt.Printf("  Blueprints: %d\n", len(env.Blueprints))
	fmt.Println("\nNext steps:")
	fmt.Println("  1. Review generated files in infra/")
	fmt.Println("  2. Run 'soloops preview' to see planned changes")
	fmt.Println("  3. Run 'soloops apply' to provision infrastructure")

	return nil
}
