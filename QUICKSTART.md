# SoloOps CLI - Quick Start Guide

Get your infrastructure up and running in 5 minutes!

## Prerequisites

1. **Install Go** (1.21 or later)
   ```bash
   # Download from https://go.dev/dl/
   # Or use a package manager:

   # macOS
   brew install go

   # Ubuntu/Debian
   sudo apt install golang-go

   # Windows
   choco install golang
   ```

2. **Install Terraform** (1.5 or later)
   ```bash
   # Download from https://www.terraform.io/downloads
   # Or use a package manager:

   # macOS
   brew install terraform

   # Ubuntu/Debian
   sudo apt install terraform

   # Windows
   choco install terraform
   ```

3. **Configure AWS Credentials**
   ```bash
   aws configure
   # Or set environment variables:
   export AWS_ACCESS_KEY_ID="your-key"
   export AWS_SECRET_ACCESS_KEY="your-secret"
   export AWS_DEFAULT_REGION="us-east-1"
   ```

## Installation

### Option 1: Build from Source

```bash
# Clone the repository
git clone https://github.com/soloops/soloops-cli.git
cd soloops-cli

# Download dependencies
go mod download

# Build the CLI
make build

# Or manually:
go build -o bin/soloops ./cmd/soloops

# Add to PATH (optional)
export PATH=$PATH:$(pwd)/bin
```

### Option 2: Install with Go

```bash
go install github.com/soloops/soloops-cli/cmd/soloops@latest
```

### Option 3: Download Binary

```bash
# Linux
curl -L https://github.com/soloops/soloops-cli/releases/latest/download/soloops-linux-amd64 -o soloops
chmod +x soloops
sudo mv soloops /usr/local/bin/

# macOS
curl -L https://github.com/soloops/soloops-cli/releases/latest/download/soloops-darwin-amd64 -o soloops
chmod +x soloops
sudo mv soloops /usr/local/bin/

# Windows
# Download from releases page and add to PATH
```

## Your First Project

### 1. Initialize a New Project

```bash
mkdir my-app
cd my-app

soloops init
```

This creates `soloops.yaml` with default configuration.

### 2. Customize Configuration

Edit `soloops.yaml`:

```yaml
project: my-awesome-app
cloud: aws
environments:
  - name: prod
    region: us-east-1
    budget_usd: 50
    blueprints:
      api:
        runtime: node18
        ingress: edge
      website:
        domain: myapp.com
policies:
  require_https: true
  deny_public_s3: true
```

### 3. Validate Configuration

```bash
soloops validate
```

Expected output:
```
✓ Configuration is valid (soloops.yaml)
  Project: my-awesome-app
  Cloud: aws
  Environments: 1
    - prod (us-east-1): $50.00 budget, 2 blueprints
```

### 4. Generate Terraform Code

```bash
soloops generate
```

This creates:
```
infra/
├── provider.tf    # Cloud provider setup
├── variables.tf   # Input variables
├── main.tf        # Infrastructure resources
├── budget.tf      # Budget alerts
└── outputs.tf     # Output values
```

### 5. Preview Changes

```bash
soloops preview
```

This runs `terraform plan` to show what will be created.

### 6. Deploy Infrastructure

```bash
soloops apply
```

Type `yes` to confirm deployment.

### 7. View Outputs

```bash
cd infra
terraform output
```

You'll see:
- API Gateway URL
- CloudFront URL
- S3 bucket name
- Lambda function ARN

## Example Workflows

### Serverless API

```bash
# 1. Initialize
soloops init

# 2. Edit soloops.yaml
cat > soloops.yaml << EOF
project: my-api
cloud: aws
environments:
  - name: prod
    region: us-east-1
    budget_usd: 20
    blueprints:
      api:
        runtime: node18
        ingress: edge
policies:
  require_https: true
EOF

# 3. Deploy
soloops validate
soloops generate
soloops apply

# 4. Deploy code
cd infra
# Copy your Lambda code
cp ../src/handler.js .
zip function.zip handler.js
aws lambda update-function-code \
  --function-name $(terraform output -raw api_lambda_arn | cut -d: -f7) \
  --zip-file fileb://function.zip

# 5. Test
curl $(terraform output -raw api_api_url)
```

### Static Website

```bash
# 1. Initialize
soloops init

# 2. Configure for static site
cat > soloops.yaml << EOF
project: my-website
cloud: aws
environments:
  - name: prod
    region: us-east-1
    budget_usd: 5
    blueprints:
      site:
        domain: mysite.com
policies:
  require_https: true
  deny_public_s3: true
EOF

# 3. Deploy infrastructure
soloops validate
soloops generate
soloops apply

# 4. Upload website
cd infra
BUCKET=$(terraform output -raw site_bucket_name)
aws s3 sync ../website/ s3://$BUCKET/

# 5. Invalidate cache
DIST_ID=$(terraform output -raw site_cloudfront_distribution_id)
aws cloudfront create-invalidation --distribution-id $DIST_ID --paths "/*"

# 6. Visit your site
open https://$(terraform output -raw site_cloudfront_url)
```

### Full-Stack Application

```yaml
project: fullstack-app
cloud: aws
environments:
  - name: prod
    region: us-east-1
    budget_usd: 100
    blueprints:
      frontend:
        domain: app.example.com
      backend:
        runtime: node18
        ingress: edge
```

```bash
soloops validate
soloops generate
soloops apply
```

## Common Commands

```bash
# Show version
soloops version

# Use different config file
soloops validate --file custom.yaml

# Target specific environment
soloops generate --env staging

# Auto-approve deployment
soloops apply --auto-approve

# Destroy infrastructure
soloops destroy
```

## Troubleshooting

### "go: command not found"

Install Go from https://go.dev/dl/

### "terraform: command not found"

Install Terraform from https://www.terraform.io/downloads

### "Error: failed to load config"

Make sure `soloops.yaml` exists in the current directory, or use `--file` flag.

### "Error: AWS credentials not configured"

Run `aws configure` or set AWS environment variables.

### "Error: validation failed"

Check the error message for details. Common issues:
- Missing required fields
- Invalid cloud provider
- Zero or negative budget
- Empty blueprints

### Terraform Errors

```bash
cd infra
terraform init
terraform validate
```

Check Terraform logs for detailed error messages.

## Next Steps

- Read [README.md](README.md) for complete documentation
- Browse [infra-templates/](infra-templates/) for blueprint examples
- Check [CONTRIBUTING.md](CONTRIBUTING.md) to contribute
- Join [Discussions](https://github.com/soloops/soloops-cli/discussions)

## Getting Help

- **Documentation**: [README.md](README.md)
- **Examples**: [infra-templates/](infra-templates/)
- **Issues**: [GitHub Issues](https://github.com/soloops/soloops-cli/issues)
- **Discussions**: [GitHub Discussions](https://github.com/soloops/soloops-cli/discussions)

## Useful Resources

- [AWS Lambda Documentation](https://docs.aws.amazon.com/lambda/)
- [Terraform Documentation](https://www.terraform.io/docs/)
- [Go Documentation](https://go.dev/doc/)
- [AWS CLI Documentation](https://aws.amazon.com/cli/)

---

**Happy Building!** 🚀