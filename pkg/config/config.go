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

package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

// Config represents the top-level soloops.yaml structure
type Config struct {
	Project      string        `yaml:"project"`
	Cloud        string        `yaml:"cloud"`
	Environments []Environment `yaml:"environments"`
	Policies     *Policies     `yaml:"policies,omitempty"`
}

// Environment represents a deployment environment
type Environment struct {
	Name       string               `yaml:"name"`
	Region     string               `yaml:"region"`
	BudgetUSD  float64              `yaml:"budget_usd"`
	Blueprints map[string]Blueprint `yaml:"blueprints"`
}

// Blueprint represents a generic infrastructure blueprint
type Blueprint struct {
	// Common fields
	Type string `yaml:"type,omitempty"`

	// Web API fields
	Runtime string `yaml:"runtime,omitempty"`
	Ingress string `yaml:"ingress,omitempty"`

	// Database fields
	DBType string `yaml:"db_type,omitempty"`

	// Static site fields
	Domain string `yaml:"domain,omitempty"`

	// Additional raw fields
	Raw map[string]interface{} `yaml:",inline"`
}

// Policies represents security and compliance policies
type Policies struct {
	RequireHTTPS bool `yaml:"require_https,omitempty"`
	DenyPublicS3 bool `yaml:"deny_public_s3,omitempty"`
}

// Load reads and parses a soloops.yaml file
func Load(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("failed to parse YAML: %w", err)
	}

	return &cfg, nil
}

// Validate checks the configuration for required fields and constraints
func (c *Config) Validate() error {
	if c.Project == "" {
		return fmt.Errorf("project name is required")
	}

	if c.Cloud == "" {
		return fmt.Errorf("cloud provider is required")
	}

	if c.Cloud != "aws" && c.Cloud != "gcp" && c.Cloud != "azure" {
		return fmt.Errorf("unsupported cloud provider: %s (supported: aws, gcp, azure)", c.Cloud)
	}

	if len(c.Environments) == 0 {
		return fmt.Errorf("at least one environment is required")
	}

	for i, env := range c.Environments {
		if env.Name == "" {
			return fmt.Errorf("environment[%d]: name is required", i)
		}
		if env.Region == "" {
			return fmt.Errorf("environment[%d] (%s): region is required", i, env.Name)
		}
		if env.BudgetUSD <= 0 {
			return fmt.Errorf("environment[%d] (%s): budget_usd must be greater than 0", i, env.Name)
		}
		if len(env.Blueprints) == 0 {
			return fmt.Errorf("environment[%d] (%s): at least one blueprint is required", i, env.Name)
		}
	}

	return nil
}

// GetEnvironment returns an environment by name
func (c *Config) GetEnvironment(name string) (*Environment, error) {
	for _, env := range c.Environments {
		if env.Name == name {
			return &env, nil
		}
	}
	return nil, fmt.Errorf("environment not found: %s", name)
}
