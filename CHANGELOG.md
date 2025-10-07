# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [0.0.1] - 2025-10-07

### Added
- Initial release of SoloOps CLI
- Core CLI commands:
  - `soloops init` - Initialize new project with soloops.yaml
  - `soloops validate` - Validate configuration with detailed error messages
  - `soloops generate` - Generate Terraform infrastructure code
  - `soloops preview` - Preview infrastructure changes (terraform plan)
  - `soloops apply` - Provision infrastructure (terraform apply)
  - `soloops destroy` - Tear down infrastructure (terraform destroy)
  - `soloops version` - Display version information
- Configuration management:
  - YAML-based configuration (soloops.yaml)
  - Schema validation with helpful error messages
  - Multi-environment support
  - Budget tracking and alerts per environment
  - Policy enforcement (HTTPS, public S3 blocking)
- AWS Blueprint support:
  - **Web API Blueprint**: Lambda + API Gateway + WAF + CloudWatch
  - **Static Site Blueprint**: S3 + CloudFront + Origin Access Identity + HTTPS
  - **Budget Alerts**: CloudWatch budgets with email notifications
- Terraform generation:
  - Clean, readable Terraform code output
  - Proper provider configuration with default tags
  - Comprehensive outputs for all resources
  - Support for AWS, GCP, and Azure (AWS fully implemented)
- Testing:
  - Unit tests for configuration validation
  - Unit tests for Terraform generation
  - Test coverage for core functionality
- Build and distribution:
  - Makefile for build automation
  - Multi-platform binary builds (Linux, macOS, Windows - amd64/arm64)
  - Docker support with multi-stage builds
  - GitHub Actions CI/CD pipeline
  - One-line installer scripts for Linux/macOS/Windows
- Documentation:
  - Comprehensive README with quick start guide
  - Contributing guidelines (CONTRIBUTING.md)
  - Testing documentation (TESTING.md)
  - Infrastructure templates and examples
  - Apache 2.0 license

### Technical Details
- Built with Go 1.21
- Dependencies:
  - cobra: CLI framework
  - gopkg.in/yaml.v3: YAML parsing
- Automated CI/CD with GitHub Actions
- golangci-lint integration for code quality

[0.0.1]: https://github.com/soloops/soloops-cli/releases/tag/v0.0.1
