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

package generator

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/soloops/soloops-cli/pkg/config"
)

// Generator generates Terraform code from SoloOps configuration
type Generator struct {
	Config *config.Config
	Env    *config.Environment
}

// New creates a new Generator instance
func New(cfg *config.Config, env *config.Environment) *Generator {
	return &Generator{
		Config: cfg,
		Env:    env,
	}
}

// Generate creates Terraform files in the infra/ directory
func (g *Generator) Generate() error {
	// Create infra directory
	if err := os.MkdirAll("infra", 0755); err != nil {
		return fmt.Errorf("failed to create infra directory: %w", err)
	}

	// Generate main files
	if err := g.generateProvider(); err != nil {
		return err
	}
	if err := g.generateVariables(); err != nil {
		return err
	}
	if err := g.generateMain(); err != nil {
		return err
	}
	if err := g.generateBudget(); err != nil {
		return err
	}
	if err := g.generateOutputs(); err != nil {
		return err
	}

	return nil
}

func (g *Generator) writeFile(filename, content string) error {
	path := filepath.Join("infra", filename)
	return os.WriteFile(path, []byte(content), 0644)
}
