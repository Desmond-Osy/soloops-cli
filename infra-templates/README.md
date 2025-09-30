# Infrastructure Templates

This directory contains blueprint templates for common infrastructure patterns.

## Available Blueprints

### 1. Serverless API

**Path**: `serverless-api/`

Creates a serverless REST API using AWS Lambda and API Gateway.

**Components**:
- AWS Lambda function
- API Gateway HTTP API
- WAF with rate limiting
- IAM roles and permissions
- CloudWatch logging

**Example Configuration**:
```yaml
blueprints:
  web_api:
    runtime: node18
    ingress: edge
```

**Documentation**: [serverless-api/README.md](serverless-api/README.md)

**Example Code**:
- Node.js: [example-handler.js](serverless-api/example-handler.js)
- Python: [example-handler.py](serverless-api/example-handler.py)

**Cost**: ~$15-20/month for 1M requests

---

### 2. Static Site

**Path**: `static-site/`

Creates a static website with S3 storage and CloudFront CDN.

**Components**:
- S3 bucket (private)
- CloudFront distribution
- Origin Access Identity
- HTTPS by default
- Custom domain support

**Example Configuration**:
```yaml
blueprints:
  static_site:
    domain: example.com
```

**Documentation**: [static-site/README.md](static-site/README.md)

**Example Website**: [static-site/example-website/](static-site/example-website/)

**Deployment Script**: [deploy.sh](static-site/deploy.sh)

**Cost**: ~$1-5/month for small-medium sites

---

### 3. Budget Alert

**Path**: `budget-alert/`

Creates AWS Budget alerts for cost monitoring.

**Components**:
- AWS Budget with monthly limits
- Email notifications at thresholds
- Cost filtering by project tags

**Configuration**:
```yaml
environments:
  - name: prod
    budget_usd: 150  # Automatically creates budget
```

**Documentation**: [budget-alert/README.md](budget-alert/README.md)

**Cost**: Free (included in AWS Free Tier)

---

## Blueprint Usage

### 1. Define in soloops.yaml

```yaml
project: my-project
cloud: aws
environments:
  - name: prod
    region: us-east-1
    budget_usd: 150
    blueprints:
      # Use multiple blueprints
      api:
        runtime: node18
        ingress: edge
      website:
        domain: mysite.com
```

### 2. Generate Infrastructure

```bash
soloops generate
```

This creates Terraform files in `infra/` directory.

### 3. Deploy

```bash
soloops apply
```

## Blueprint Development

### Creating a New Blueprint

1. Create a directory: `infra-templates/my-blueprint/`
2. Add README.md with documentation
3. Add example code/configuration
4. Update generator in `pkg/generator/main.go`
5. Add blueprint type to `pkg/config/config.go`
6. Write tests in `tests/generator_test.go`

### Blueprint Structure

```
my-blueprint/
├── README.md              # Documentation
├── example-code/          # Sample implementations
├── terraform/             # Optional: Terraform modules
└── deploy.sh              # Optional: Deployment script
```

## Cloud Provider Support

### AWS (Full Support)
- ✅ Serverless API
- ✅ Static Site
- ✅ Budget Alert

### GCP (Planned)
- 🚧 Cloud Functions + API Gateway
- 🚧 Cloud Storage + Cloud CDN
- 🚧 Budget Alerts

### Azure (Planned)
- 🚧 Azure Functions + API Management
- 🚧 Blob Storage + Azure CDN
- 🚧 Cost Management Alerts

## Common Patterns

### Full-Stack Application

```yaml
blueprints:
  frontend:
    type: static_site
    domain: app.example.com
  backend:
    type: web_api
    runtime: node18
  database:
    type: aurora_serverless_v2
```

### Microservices

```yaml
blueprints:
  user_service:
    runtime: go1.x
    ingress: regional
  order_service:
    runtime: python311
    ingress: regional
  notification_service:
    runtime: node20
    ingress: regional
```

### Multi-Region

```yaml
environments:
  - name: us-prod
    region: us-east-1
    budget_usd: 300
    blueprints:
      api:
        runtime: node18
  - name: eu-prod
    region: eu-west-1
    budget_usd: 300
    blueprints:
      api:
        runtime: node18
```

## Best Practices

### 1. Tagging
All resources are automatically tagged with:
- Project name
- Environment
- ManagedBy: "SoloOps"

### 2. Security
- HTTPS enforced by default
- Private resources (no public access)
- IAM least privilege
- WAF rate limiting
- CloudWatch logging

### 3. Cost Optimization
- Serverless architecture (pay per use)
- CDN caching
- Budget alerts
- Auto-scaling

### 4. Monitoring
- CloudWatch metrics
- Access logs
- Error tracking
- Performance monitoring

## Template Customization

After generating Terraform files, you can:

1. **Add resources** directly to generated `.tf` files
2. **Modify configurations** to fit your needs
3. **Add variables** in `variables.tf`
4. **Customize outputs** in `outputs.tf`

SoloOps generates a starting point - customize as needed!

## Contributing Templates

We welcome new blueprint contributions!

See [CONTRIBUTING.md](../CONTRIBUTING.md) for guidelines on:
- Template structure
- Documentation requirements
- Testing requirements
- Code review process

## Support

- Documentation: [Main README](../README.md)
- Issues: [GitHub Issues](https://github.com/soloops/soloops-cli/issues)
- Discussions: [GitHub Discussions](https://github.com/soloops/soloops-cli/discussions)

## License

All templates are licensed under Apache 2.0 - see [LICENSE](../LICENSE)