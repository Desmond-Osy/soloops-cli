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

// DefaultTemplate returns a starter soloops.yaml template
func DefaultTemplate() string {
	return `project: my-project
cloud: aws
environments:
  - name: prod
    region: us-east-1
    budget_usd: 150
    blueprints:
      web_api:
        runtime: node18
        ingress: edge
      static_site:
        domain: example.com
policies:
  require_https: true
  deny_public_s3: true
`
}