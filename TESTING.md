# Testing SoloOps CLI

This document describes how to test the SoloOps CLI after installation.

## Prerequisites

1. Install Go 1.21 or later: https://go.dev/dl/
2. Install Terraform: https://www.terraform.io/downloads

## Building the CLI

```bash
# Install dependencies
go mod download

# Build the binary
go build -o bin/soloops ./cmd/soloops

# Or use Make
make build
```

## Running Tests

```bash
# Run all unit tests
go test -v ./tests/

# Run tests with coverage
go test -v -coverprofile=coverage.out ./tests/
go tool cover -html=coverage.out

# Or use Make
make test
make test-coverage
```

## Manual Testing

### 1. Test `soloops version`

```bash
./bin/soloops version
```

Expected output:
```
SoloOps CLI
  Version:    dev
  Git Commit: none
  Build Date: unknown
```

### 2. Test `soloops init`

```bash
# Create a test directory
mkdir test-project
cd test-project

# Initialize a new project
../bin/soloops init
```

Expected:
- Creates `soloops.yaml` file
- Shows next steps instructions

### 3. Test `soloops validate`

```bash
# Validate the generated config
../bin/soloops validate
```

Expected output:
```
✓ Configuration is valid (soloops.yaml)
  Project: my-project
  Cloud: aws
  Environments: 1
    - prod (us-east-1): $150.00 budget, 2 blueprints
```

### 4. Test with invalid config

Edit `soloops.yaml` and remove the `project` field, then:

```bash
../bin/soloops validate
```

Expected: Error message about missing project name

### 5. Test `soloops generate`

Restore valid config and run:

```bash
../bin/soloops generate
```

Expected:
- Creates `infra/` directory
- Generates Terraform files:
  - `provider.tf`
  - `variables.tf`
  - `main.tf`
  - `budget.tf`
  - `outputs.tf`

### 6. Verify generated Terraform

```bash
cd infra
terraform init
terraform validate
terraform plan
```

Expected:
- Terraform validates successfully
- Plan shows resources to be created

### 7. Test `soloops preview`

```bash
cd ..
../bin/soloops preview
```

Expected:
- Runs `terraform init`
- Runs `terraform plan`
- Shows planned infrastructure changes

### 8. Test help commands

```bash
../bin/soloops --help
../bin/soloops init --help
../bin/soloops validate --help
```

Expected: Shows usage information for each command

## Integration Tests

### Test different cloud providers

Edit `soloops.yaml` and change `cloud: aws` to:
- `cloud: gcp`
- `cloud: azure`

Run `soloops generate` for each and verify correct provider files are generated.

### Test different blueprints

Modify the `blueprints` section to test different configurations:

```yaml
blueprints:
  # Just web API
  web_api:
    runtime: python39
    ingress: edge
```

```yaml
blueprints:
  # Just static site
  static_site:
    domain: test.example.com
```

Run `soloops generate` and verify correct resources are created.

### Test multiple environments

Edit `soloops.yaml` to add multiple environments:

```yaml
environments:
  - name: dev
    region: us-west-2
    budget_usd: 50
    blueprints:
      web_api:
        runtime: node18
        ingress: edge
  - name: prod
    region: us-east-1
    budget_usd: 150
    blueprints:
      web_api:
        runtime: node18
        ingress: edge
```

Test generating for specific environment:

```bash
../bin/soloops generate --env dev
../bin/soloops generate --env prod
```

## Automated Test Suite

Run the full test suite:

```bash
# From project root
make test

# With verbose output
go test -v ./tests/

# Specific test
go test -v -run TestConfigValidate ./tests/
go test -v -run TestGeneratorGenerate ./tests/
```

Expected: All tests pass

## Docker Testing

Build and test in Docker:

```bash
# Build Docker image
docker build -t soloops:test .

# Test version command
docker run --rm soloops:test version

# Test init command with volume mount
mkdir docker-test
docker run --rm -v $(pwd)/docker-test:/workspace soloops:test init
cat docker-test/soloops.yaml
```

## Common Issues

### Go not found
Install Go from https://go.dev/dl/

### Terraform not found
Install Terraform from https://www.terraform.io/downloads

### Module errors
Run `go mod download` to fetch dependencies

### Permission errors
Ensure the binary has execute permissions: `chmod +x bin/soloops`

## Success Criteria

The CLI is working correctly if:

1. ✅ `soloops init` creates a valid soloops.yaml
2. ✅ `soloops validate` correctly validates and reports errors
3. ✅ `soloops generate` creates valid Terraform files
4. ✅ Generated Terraform passes `terraform validate`
5. ✅ All unit tests pass
6. ✅ Help commands display correctly
7. ✅ Version command shows build info

## Next Steps

After successful testing:

1. Run `make build-all` to create binaries for all platforms
2. Test on different operating systems
3. Create GitHub release with binaries
4. Update documentation with any findings