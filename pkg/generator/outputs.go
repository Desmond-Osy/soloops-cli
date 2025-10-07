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
	"strings"
)

func (g *Generator) generateOutputs() error {
	var outputs strings.Builder

	outputs.WriteString("# Terraform outputs\n\n")

	// Generate outputs for each blueprint
	for name, blueprint := range g.Env.Blueprints {
		if blueprint.Runtime != "" || blueprint.Ingress != "" {
			// Web API outputs
			outputs.WriteString(fmt.Sprintf(`output "%s_api_url" {
  description = "API Gateway endpoint URL for %s"
  value       = try(aws_apigatewayv2_stage.%s.invoke_url, "N/A")
}

output "%s_lambda_arn" {
  description = "Lambda function ARN for %s"
  value       = try(aws_lambda_function.%s.arn, "N/A")
}

`, name, name, name, name, name, name))
		}

		if blueprint.Domain != "" {
			// Static site outputs
			outputs.WriteString(fmt.Sprintf(`output "%s_bucket_name" {
  description = "S3 bucket name for %s"
  value       = try(aws_s3_bucket.%s.id, "N/A")
}

output "%s_cloudfront_url" {
  description = "CloudFront distribution URL for %s"
  value       = try(aws_cloudfront_distribution.%s.domain_name, "N/A")
}

`, name, name, name, name, name, name))
		}
	}

	outputs.WriteString(`output "environment" {
  description = "Environment name"
  value       = var.environment
}

output "region" {
  description = "Deployment region"
  value       = var.region
}
`)

	return g.writeFile("outputs.tf", outputs.String())
}
