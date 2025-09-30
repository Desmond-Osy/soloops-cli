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

func (g *Generator) generateVariables() error {
	content := fmt.Sprintf(`variable "project_name" {
  description = "Project name"
  type        = string
  default     = "%s"
}

variable "environment" {
  description = "Environment name"
  type        = string
  default     = "%s"
}

variable "region" {
  description = "Cloud region"
  type        = string
  default     = "%s"
}

variable "budget_usd" {
  description = "Monthly budget in USD"
  type        = number
  default     = %.2f
}
`, g.Config.Project, g.Env.Name, g.Env.Region, g.Env.BudgetUSD)

	return g.writeFile("variables.tf", content)
}