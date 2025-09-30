# SoloOps CLI

[![CI](https://github.com/soloops/soloops-cli/workflows/CI/badge.svg)](https://github.com/soloops/soloops-cli/actions)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/soloops/soloops-cli)](https://goreportcard.com/report/github.com/soloops/soloops-cli)

SoloOps is a command-line tool for scaffolding, validating, and managing infrastructure blueprints described in a YAML manifest (`soloops.yaml`). It generates Terraform code from your declarative configuration, making it easy to provision cloud resources with best practices built-in.

## Features

- **Declarative Infrastructure**: Define your infrastructure in a simple YAML manifest
- **Blueprint System**: Pre-built templates for common patterns (serverless APIs, static sites, databases)
- **Multi-Cloud Support**: AWS, GCP, and Azure (AWS fully implemented in MVP)
- **Budget Aware**: Automatic budget alerts and cost controls
- **Security First**: Built-in WAF, HTTPS enforcement, and compliance policies
- **Terraform Generation**: Generates clean, readable Terraform code
- **Easy to Use**: Simple CLI commands for the entire lifecycle

## Quick Start

### Installation

#### Using pre-built binaries

Download the latest release for your platform from the [releases page](https://github.com/soloops/soloops-cli/releases).

```bash
# Linux/macOS
curl -L https://github.com/soloops/soloops-cli/releases/latest/download/soloops-$(uname -s | tr '[:upper:]' '[:lower:]')-$(uname -m) -o soloops
chmod +x soloops
sudo mv soloops /usr/local/bin/

# Verify installation
soloops version
```

#### Using Go

```bash
go install github.com/soloops/soloops-cli/cmd/soloops@latest
```

#### Using Docker

```bash
docker pull soloops/soloops-cli:latest
docker run --rm -v $(pwd):/workspace soloops/soloops-cli:latest init
```

### Basic Usage

1. **Initialize a new project**:

```bash
soloops init
```

This creates a `soloops.yaml` manifest with sensible defaults.

2. **Customize your configuration**:

Edit `soloops.yaml` to define your infrastructure:

```yaml
project: my-awesome-app
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
        domain: myapp.com
policies:
  require_https: true
  deny_public_s3: true
```

3. **Validate your configuration**:

```bash
soloops validate
```

4. **Generate Terraform code**:

```bash
soloops generate
```

This creates Terraform files in the `infra/` directory.

5. **Preview changes**:

```bash
soloops preview
```

Shows what infrastructure will be created (runs `terraform plan`).

6. **Apply changes**:

```bash
soloops apply
```

Provisions your infrastructure (runs `terraform apply`).

7. **Destroy when done**:

```bash
soloops destroy
```

## Commands

| Command | Description |
|---------|-------------|
| `soloops init` | Create a new soloops.yaml manifest |
| `soloops validate` | Validate the configuration |
| `soloops generate` | Generate Terraform files |
| `soloops preview` | Preview infrastructure changes |
| `soloops apply` | Provision infrastructure |
| `soloops destroy` | Destroy infrastructure |
| `soloops version` | Show version information |

### Global Flags

- `--file, -f`: Path to soloops.yaml (default: `soloops.yaml`)
- `--env, -e`: Target environment (defaults to first in manifest)

## Configuration

### Project Structure

```
my-project/
├── soloops.yaml          # Your infrastructure manifest
├── infra/                # Generated Terraform files
│   ├── provider.tf
│   ├── variables.tf
│   ├── main.tf
│   ├── budget.tf
│   └── outputs.tf
└── terraform.tfstate     # Terraform state (created after apply)
```

### Example soloops.yaml

```yaml
project: acme-api
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
        domain: acme.com
      database:
        type: aurora_serverless_v2
policies:
  require_https: true
  deny_public_s3: true
```

## Supported Blueprints

### Web API (AWS)

Creates a serverless API with:
- AWS Lambda function
- API Gateway HTTP API
- WAF with rate limiting
- CloudWatch logs

```yaml
web_api:
  runtime: node18
  ingress: edge
```

### Static Site (AWS)

Creates a static website with:
- S3 bucket
- CloudFront distribution
- HTTPS by default
- Origin access identity

```yaml
static_site:
  domain: example.com
```

### Database (Coming Soon)

Support for managed databases:
- Aurora Serverless v2
- RDS instances
- DynamoDB tables

## Development

### Prerequisites

- Go 1.21 or later
- Make
- Docker (optional)

### Building from Source

```bash
# Clone the repository
git clone https://github.com/soloops/soloops-cli.git
cd soloops-cli

# Download dependencies
make deps

# Build the binary
make build

# Run tests
make test

# Install locally
make install
```

### Build Commands

```bash
make build          # Build the CLI binary
make build-all      # Build binaries for all platforms
make test           # Run tests
make test-coverage  # Run tests with coverage report
make install        # Install the CLI locally
make clean          # Clean build artifacts
make lint           # Run linters
make format         # Format code
make docker-build   # Build Docker image
```

## Contributing

We welcome contributions! Please see [CONTRIBUTING.md](CONTRIBUTING.md) for guidelines.

### Code of Conduct

This project adheres to a [Code of Conduct](CODE_OF_CONDUCT.md). By participating, you are expected to uphold this code.

### Security

For security issues, please see [SECURITY.md](SECURITY.md) for our security policy and how to report vulnerabilities.

## Roadmap

- [ ] Multi-environment support
- [ ] Remote state backends (S3, GCS, Azure Blob)
- [ ] Cost estimation integration (Infracost)
- [ ] Kubernetes blueprint support
- [ ] Database blueprint implementation
- [ ] Custom blueprint plugins
- [ ] Interactive CLI mode
- [ ] Drift detection
- [ ] Policy-as-code validation (OPA)

## License

This project is licensed under the Apache License 2.0 - see the [LICENSE](LICENSE) file for details.

## Support

- Documentation: [https://soloops.dev/docs](https://soloops.dev/docs)
- Issues: [GitHub Issues](https://github.com/soloops/soloops-cli/issues)
- Discussions: [GitHub Discussions](https://github.com/soloops/soloops-cli/discussions)

## Acknowledgments

SoloOps is inspired by and builds upon the excellent work of:
- [Terraform](https://www.terraform.io/) by HashiCorp
- [Cobra](https://github.com/spf13/cobra) CLI framework
- The open-source infrastructure-as-code community

---

Built with ❤️ by the SoloOps community