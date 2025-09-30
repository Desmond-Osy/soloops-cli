# Static Site Blueprint

This blueprint creates a static website hosted on AWS S3 with CloudFront CDN.

## Architecture

```
Internet → CloudFront (HTTPS) → S3 Bucket (Private)
           ↓
       Origin Access Identity
```

## Components

### S3 Bucket
- Private bucket (no public access)
- Website hosting configuration
- Versioning enabled
- Server-side encryption

### CloudFront Distribution
- Global CDN for fast content delivery
- HTTPS enforcement
- Custom domain support (optional)
- Automatic compression
- Cache optimization

### Origin Access Identity (OAI)
- Secure S3 access from CloudFront only
- No direct public S3 access
- IAM-based permissions

## Configuration

```yaml
blueprints:
  static_site:
    domain: example.com     # Your domain (optional)
    certificate_arn: ""     # ACM certificate ARN (optional)
```

## Generated Resources

1. **S3 Bucket** (`aws_s3_bucket`)
   - Private bucket for static files
   - Website configuration

2. **S3 Bucket Public Access Block** (`aws_s3_bucket_public_access_block`)
   - Blocks all public access
   - Enforces security best practices

3. **CloudFront Distribution** (`aws_cloudfront_distribution`)
   - CDN distribution
   - HTTPS redirect
   - Gzip compression

4. **Origin Access Identity** (`aws_cloudfront_origin_access_identity`)
   - Secure CloudFront → S3 access

5. **S3 Bucket Policy** (`aws_s3_bucket_policy`)
   - Allows CloudFront OAI to read objects
   - Denies other access

## Outputs

- `{name}_bucket_name` - S3 bucket name
- `{name}_cloudfront_url` - CloudFront distribution URL

## Deployment

### 1. Generate Infrastructure

```bash
soloops generate
```

### 2. Deploy Infrastructure

```bash
cd infra
terraform init
terraform apply
```

### 3. Upload Your Website

```bash
# Upload files to S3
aws s3 sync ./website/ s3://YOUR_BUCKET_NAME/ --delete

# Invalidate CloudFront cache
aws cloudfront create-invalidation \
  --distribution-id YOUR_DISTRIBUTION_ID \
  --paths "/*"
```

### 4. Access Your Site

Visit the CloudFront URL from the Terraform outputs.

## Cost Estimate

Based on moderate traffic (AWS us-east-1):

- **S3 Storage**: $0.023 per GB/month
  - Example: 1GB website = $0.023/month
- **S3 Requests**: $0.0004 per 1000 GET requests
  - Example: 100K requests/month = $0.04/month
- **CloudFront**: $0.085 per GB transferred
  - Example: 10GB/month = $0.85/month
- **CloudFront Requests**: $0.0075 per 10,000 requests
  - Example: 100K requests/month = $0.075/month

**Estimated Total**: $1-5/month for small-medium sites

## Security Features

- ✅ Private S3 bucket (no public access)
- ✅ HTTPS only (redirects HTTP to HTTPS)
- ✅ Origin Access Identity for secure access
- ✅ Server-side encryption
- ✅ CloudFront logging (optional)

## Example Website Structure

```
website/
├── index.html
├── error.html
├── css/
│   └── style.css
├── js/
│   └── main.js
└── images/
    └── logo.png
```

### Sample index.html

```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>My Static Site</title>
    <link rel="stylesheet" href="css/style.css">
</head>
<body>
    <header>
        <h1>Welcome to My Site</h1>
    </header>
    <main>
        <p>This site is hosted on S3 with CloudFront!</p>
    </main>
    <footer>
        <p>&copy; 2025 My Company</p>
    </footer>
    <script src="js/main.js"></script>
</body>
</html>
```

### Sample error.html

```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Page Not Found</title>
</head>
<body>
    <h1>404 - Page Not Found</h1>
    <p>Sorry, the page you're looking for doesn't exist.</p>
    <a href="/">Go back home</a>
</body>
</html>
```

## Custom Domain Setup

### 1. Request SSL Certificate

```bash
aws acm request-certificate \
  --domain-name example.com \
  --subject-alternative-names www.example.com \
  --validation-method DNS \
  --region us-east-1
```

Note: Certificate must be in us-east-1 for CloudFront.

### 2. Update soloops.yaml

```yaml
blueprints:
  static_site:
    domain: example.com
    certificate_arn: arn:aws:acm:us-east-1:123456789012:certificate/abc123
```

### 3. Add to Terraform

After generation, add to CloudFront distribution:

```hcl
resource "aws_cloudfront_distribution" "static_site" {
  # ... existing config ...

  aliases = ["example.com", "www.example.com"]

  viewer_certificate {
    acm_certificate_arn      = var.certificate_arn
    ssl_support_method       = "sni-only"
    minimum_protocol_version = "TLSv1.2_2021"
  }
}
```

### 4. Update DNS

Point your domain to CloudFront:

```
example.com     CNAME  d123456abcdef.cloudfront.net
www.example.com CNAME  d123456abcdef.cloudfront.net
```

## Advanced Configuration

### Enable Access Logs

Add to S3 bucket:

```hcl
resource "aws_s3_bucket_logging" "static_site" {
  bucket = aws_s3_bucket.static_site.id

  target_bucket = aws_s3_bucket.logs.id
  target_prefix = "access-logs/"
}
```

### Add Cache Behaviors

Customize caching for different paths:

```hcl
resource "aws_cloudfront_distribution" "static_site" {
  # ... existing config ...

  ordered_cache_behavior {
    path_pattern     = "/images/*"
    allowed_methods  = ["GET", "HEAD"]
    cached_methods   = ["GET", "HEAD"]
    target_origin_id = "S3-${aws_s3_bucket.static_site.id}"

    forwarded_values {
      query_string = false
      cookies {
        forward = "none"
      }
    }

    min_ttl                = 86400
    default_ttl            = 86400
    max_ttl                = 31536000
    viewer_protocol_policy = "redirect-to-https"
  }
}
```

### Add Lambda@Edge Functions

Add edge functions for dynamic behavior:

```hcl
resource "aws_lambda_function" "edge_function" {
  provider      = aws.us-east-1
  filename      = "edge-function.zip"
  function_name = "cloudfront-edge-function"
  role          = aws_iam_role.edge_function.arn
  handler       = "index.handler"
  runtime       = "nodejs18.x"
  publish       = true
}

resource "aws_cloudfront_distribution" "static_site" {
  # ... existing config ...

  default_cache_behavior {
    # ... existing config ...

    lambda_function_association {
      event_type   = "viewer-request"
      lambda_arn   = aws_lambda_function.edge_function.qualified_arn
    }
  }
}
```

## Deployment Automation

### GitHub Actions Workflow

```yaml
name: Deploy Static Site

on:
  push:
    branches: [ main ]

jobs:
  deploy:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4

      - name: Configure AWS credentials
        uses: aws-actions/configure-aws-credentials@v4
        with:
          aws-access-key-id: ${{ secrets.AWS_ACCESS_KEY_ID }}
          aws-secret-access-key: ${{ secrets.AWS_SECRET_ACCESS_KEY }}
          aws-region: us-east-1

      - name: Sync to S3
        run: |
          aws s3 sync ./website s3://${{ secrets.S3_BUCKET }}/ --delete

      - name: Invalidate CloudFront
        run: |
          aws cloudfront create-invalidation \
            --distribution-id ${{ secrets.CLOUDFRONT_ID }} \
            --paths "/*"
```

## Monitoring

### CloudWatch Metrics
- Requests
- Bytes downloaded
- Error rate (4xx, 5xx)
- Cache hit ratio

### CloudFront Reports
- Popular objects
- Top referrers
- Usage by location
- Viewer details

## Troubleshooting

### 403 Forbidden Errors
- Check S3 bucket policy allows CloudFront OAI
- Verify bucket doesn't have public access block

### Stale Content
- Invalidate CloudFront cache
- Check cache behavior TTL settings

### Custom Domain Not Working
- Verify certificate is in us-east-1
- Check DNS CNAME records
- Wait for CloudFront distribution to deploy

## Best Practices

1. **Use CloudFront Functions** for lightweight edge logic
2. **Enable compression** for text-based files
3. **Set appropriate cache TTLs** for different content types
4. **Use versioned filenames** for cache busting (e.g., `app.v2.js`)
5. **Enable logging** for troubleshooting
6. **Set up CloudWatch alarms** for error rates

## References

- [S3 Static Website Hosting](https://docs.aws.amazon.com/AmazonS3/latest/userguide/WebsiteHosting.html)
- [CloudFront Documentation](https://docs.aws.amazon.com/cloudfront/)
- [CloudFront Security](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/security.html)