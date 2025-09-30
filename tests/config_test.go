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

package tests

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/soloops/soloops-cli/pkg/config"
)

func TestConfigLoad(t *testing.T) {
	// Create a temporary test file
	tmpDir := t.TempDir()
	configPath := filepath.Join(tmpDir, "soloops.yaml")

	validConfig := `project: test-project
cloud: aws
environments:
  - name: prod
    region: us-east-1
    budget_usd: 100
    blueprints:
      web_api:
        runtime: node18
        ingress: edge
`

	if err := os.WriteFile(configPath, []byte(validConfig), 0644); err != nil {
		t.Fatalf("Failed to write test config: %v", err)
	}

	// Test loading
	cfg, err := config.Load(configPath)
	if err != nil {
		t.Fatalf("Failed to load config: %v", err)
	}

	if cfg.Project != "test-project" {
		t.Errorf("Expected project 'test-project', got '%s'", cfg.Project)
	}

	if cfg.Cloud != "aws" {
		t.Errorf("Expected cloud 'aws', got '%s'", cfg.Cloud)
	}

	if len(cfg.Environments) != 1 {
		t.Errorf("Expected 1 environment, got %d", len(cfg.Environments))
	}
}

func TestConfigValidate(t *testing.T) {
	tests := []struct {
		name        string
		config      *config.Config
		expectError bool
	}{
		{
			name: "valid config",
			config: &config.Config{
				Project: "test",
				Cloud:   "aws",
				Environments: []config.Environment{
					{
						Name:      "prod",
						Region:    "us-east-1",
						BudgetUSD: 100,
						Blueprints: map[string]config.Blueprint{
							"api": {},
						},
					},
				},
			},
			expectError: false,
		},
		{
			name: "missing project",
			config: &config.Config{
				Cloud: "aws",
				Environments: []config.Environment{
					{
						Name:      "prod",
						Region:    "us-east-1",
						BudgetUSD: 100,
						Blueprints: map[string]config.Blueprint{
							"api": {},
						},
					},
				},
			},
			expectError: true,
		},
		{
			name: "invalid cloud",
			config: &config.Config{
				Project: "test",
				Cloud:   "invalid",
				Environments: []config.Environment{
					{
						Name:      "prod",
						Region:    "us-east-1",
						BudgetUSD: 100,
						Blueprints: map[string]config.Blueprint{
							"api": {},
						},
					},
				},
			},
			expectError: true,
		},
		{
			name: "no environments",
			config: &config.Config{
				Project:      "test",
				Cloud:        "aws",
				Environments: []config.Environment{},
			},
			expectError: true,
		},
		{
			name: "zero budget",
			config: &config.Config{
				Project: "test",
				Cloud:   "aws",
				Environments: []config.Environment{
					{
						Name:      "prod",
						Region:    "us-east-1",
						BudgetUSD: 0,
						Blueprints: map[string]config.Blueprint{
							"api": {},
						},
					},
				},
			},
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.config.Validate()
			if tt.expectError && err == nil {
				t.Error("Expected error but got none")
			}
			if !tt.expectError && err != nil {
				t.Errorf("Expected no error but got: %v", err)
			}
		})
	}
}

func TestGetEnvironment(t *testing.T) {
	cfg := &config.Config{
		Project: "test",
		Cloud:   "aws",
		Environments: []config.Environment{
			{Name: "dev", Region: "us-west-1", BudgetUSD: 50},
			{Name: "prod", Region: "us-east-1", BudgetUSD: 100},
		},
	}

	// Test existing environment
	env, err := cfg.GetEnvironment("prod")
	if err != nil {
		t.Fatalf("Failed to get environment: %v", err)
	}
	if env.Name != "prod" {
		t.Errorf("Expected environment 'prod', got '%s'", env.Name)
	}

	// Test non-existent environment
	_, err = cfg.GetEnvironment("staging")
	if err == nil {
		t.Error("Expected error for non-existent environment")
	}
}