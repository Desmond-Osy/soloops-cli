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

import "fmt"

func (g *Generator) generateProvider() error {
	var content string

	switch g.Config.Cloud {
	case "aws":
		content = fmt.Sprintf(`terraform {
  required_version = ">= 1.5"

  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 5.0"
    }
  }
}

provider "aws" {
  region = "%s"

  default_tags {
    tags = {
      Project     = "%s"
      Environment = "%s"
      ManagedBy   = "SoloOps"
    }
  }
}
`, g.Env.Region, g.Config.Project, g.Env.Name)

	case "gcp":
		content = fmt.Sprintf(`terraform {
  required_version = ">= 1.5"

  required_providers {
    google = {
      source  = "hashicorp/google"
      version = "~> 5.0"
    }
  }
}

provider "google" {
  project = "%s"
  region  = "%s"

  default_labels = {
    project     = "%s"
    environment = "%s"
    managed_by  = "soloops"
  }
}
`, g.Config.Project, g.Env.Region, g.Config.Project, g.Env.Name)

	case "azure":
		content = fmt.Sprintf(`terraform {
  required_version = ">= 1.5"

  required_providers {
    azurerm = {
      source  = "hashicorp/azurerm"
      version = "~> 3.0"
    }
  }
}

provider "azurerm" {
  features {}

  tags = {
    Project     = "%s"
    Environment = "%s"
    ManagedBy   = "SoloOps"
  }
}
`, g.Config.Project, g.Env.Name)

	default:
		return fmt.Errorf("unsupported cloud provider: %s", g.Config.Cloud)
	}

	return g.writeFile("provider.tf", content)
}
