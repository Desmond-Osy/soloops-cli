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
	"github.com/spf13/cobra"
)

var (
	configFile string
	envName    string
	version    string
	gitCommit  string
	buildDate  string
)

// SetVersionInfo sets version information from build-time variables
func SetVersionInfo(v, commit, date string) {
	version = v
	gitCommit = commit
	buildDate = date
}

var rootCmd = &cobra.Command{
	Use:   "soloops",
	Short: "SoloOps - Infrastructure blueprint management CLI",
	Long: `SoloOps is a command-line tool for scaffolding, validating, and managing
infrastructure blueprints described in a YAML manifest (soloops.yaml).

It generates Terraform code from your declarative configuration, making it
easy to provision cloud resources with best practices built-in.`,
}

// Execute runs the root command
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&configFile, "file", "f", "soloops.yaml", "Path to soloops.yaml manifest")
	rootCmd.PersistentFlags().StringVarP(&envName, "env", "e", "", "Environment to target (defaults to first in manifest)")

	// Add subcommands
	rootCmd.AddCommand(initCmd)
	rootCmd.AddCommand(validateCmd)
	rootCmd.AddCommand(generateCmd)
	rootCmd.AddCommand(previewCmd)
	rootCmd.AddCommand(applyCmd)
	rootCmd.AddCommand(destroyCmd)
	rootCmd.AddCommand(versionCmd)
}
