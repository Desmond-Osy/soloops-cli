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
	"strings"
	"testing"

	"github.com/soloops/soloops-cli/pkg/config"
	"github.com/soloops/soloops-cli/pkg/generator"
)

func TestGeneratorGenerate(t *testing.T) {
	// Create a temporary directory for generated files
	tmpDir := t.TempDir()
	originalDir, _ := os.Getwd()
	defer func() {
		if err := os.Chdir(originalDir); err != nil {
			t.Errorf("Failed to change back to original directory: %v", err)
		}
	}()
	if err := os.Chdir(tmpDir); err != nil {
		t.Fatalf("Failed to change to temp directory: %v", err)
	}

	cfg := &config.Config{
		Project: "test-project",
		Cloud:   "aws",
		Policies: &config.Policies{
			RequireHTTPS: true,
			DenyPublicS3: true,
		},
	}

	env := &config.Environment{
		Name:      "prod",
		Region:    "us-east-1",
		BudgetUSD: 150,
		Blueprints: map[string]config.Blueprint{
			"web_api": {
				Runtime: "node18",
				Ingress: "edge",
			},
			"static_site": {
				Domain: "example.com",
			},
		},
	}

	gen := generator.New(cfg, env)
	if err := gen.Generate(); err != nil {
		t.Fatalf("Failed to generate: %v", err)
	}

	// Check that expected files were created
	expectedFiles := []string{
		"infra/provider.tf",
		"infra/variables.tf",
		"infra/main.tf",
		"infra/budget.tf",
		"infra/outputs.tf",
	}

	for _, file := range expectedFiles {
		if _, err := os.Stat(file); os.IsNotExist(err) {
			t.Errorf("Expected file not created: %s", file)
		}
	}

	// Check provider.tf content
	providerContent, err := os.ReadFile("infra/provider.tf")
	if err != nil {
		t.Fatalf("Failed to read provider.tf: %v", err)
	}

	if !strings.Contains(string(providerContent), "hashicorp/aws") {
		t.Error("provider.tf should contain AWS provider")
	}

	if !strings.Contains(string(providerContent), "us-east-1") {
		t.Error("provider.tf should contain region")
	}

	// Check main.tf content
	mainContent, err := os.ReadFile("infra/main.tf")
	if err != nil {
		t.Fatalf("Failed to read main.tf: %v", err)
	}

	if !strings.Contains(string(mainContent), "aws_lambda_function") {
		t.Error("main.tf should contain Lambda function for web_api")
	}

	if !strings.Contains(string(mainContent), "aws_s3_bucket") {
		t.Error("main.tf should contain S3 bucket for static_site")
	}

	if !strings.Contains(string(mainContent), "aws_cloudfront_distribution") {
		t.Error("main.tf should contain CloudFront distribution for static_site")
	}

	// Check budget.tf content
	budgetContent, err := os.ReadFile("infra/budget.tf")
	if err != nil {
		t.Fatalf("Failed to read budget.tf: %v", err)
	}

	if !strings.Contains(string(budgetContent), "150.00") {
		t.Error("budget.tf should contain budget amount")
	}
}

func TestGeneratorDifferentClouds(t *testing.T) {
	tmpDir := t.TempDir()
	originalDir, _ := os.Getwd()
	defer func() {
		if err := os.Chdir(originalDir); err != nil {
			t.Errorf("Failed to change back to original directory: %v", err)
		}
	}()

	tests := []struct {
		cloud            string
		expectedProvider string
	}{
		{"aws", "hashicorp/aws"},
		{"gcp", "hashicorp/google"},
		{"azure", "hashicorp/azurerm"},
	}

	for _, tt := range tests {
		t.Run(tt.cloud, func(t *testing.T) {
			cloudDir := filepath.Join(tmpDir, tt.cloud)
			if err := os.MkdirAll(cloudDir, 0755); err != nil {
				t.Fatalf("Failed to create cloud directory: %v", err)
			}
			if err := os.Chdir(cloudDir); err != nil {
				t.Fatalf("Failed to change to cloud directory: %v", err)
			}

			cfg := &config.Config{
				Project: "test",
				Cloud:   tt.cloud,
			}

			env := &config.Environment{
				Name:      "prod",
				Region:    "us-east-1",
				BudgetUSD: 100,
				Blueprints: map[string]config.Blueprint{
					"test": {},
				},
			}

			gen := generator.New(cfg, env)
			if err := gen.Generate(); err != nil {
				t.Fatalf("Failed to generate for %s: %v", tt.cloud, err)
			}

			content, err := os.ReadFile("infra/provider.tf")
			if err != nil {
				t.Fatalf("Failed to read provider.tf: %v", err)
			}

			if !strings.Contains(string(content), tt.expectedProvider) {
				t.Errorf("Expected provider %s not found in provider.tf", tt.expectedProvider)
			}
		})
	}
}
