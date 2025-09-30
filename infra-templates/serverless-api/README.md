# Serverless API Blueprint

This blueprint creates a serverless API using AWS Lambda and API Gateway.

## Architecture

```
Internet → API Gateway → Lambda Function → (Optional) Database
           ↓
          WAF (Rate Limiting)
```

## Components

### AWS Lambda
- Serverless compute for API logic
- Configurable runtime (Node.js, Python, Go, etc.)
- Auto-scaling
- Pay per invocation

### API Gateway HTTP API
- RESTful API endpoint
- Edge-optimized or regional
- Custom domain support
- Request/response transformation

### WAF (Web Application Firewall)
- Rate limiting (2000 requests per 5 minutes per IP)
- DDoS protection
- IP filtering (optional)
- Custom rules support

### IAM Roles
- Least-privilege execution role
- CloudWatch Logs permissions
- Additional permissions as needed

## Configuration

```yaml
blueprints:
  web_api:
    runtime: node18        # Lambda runtime
    ingress: edge          # edge or regional
    memory: 512            # Optional: Memory in MB (default: 128)
    timeout: 30            # Optional: Timeout in seconds (default: 3)
```

## Supported Runtimes

- `node18` - Node.js 18.x
- `node20` - Node.js 20.x
- `python39` - Python 3.9
- `python310` - Python 3.10
- `python311` - Python 3.11
- `go1.x` - Go 1.x
- `java17` - Java 17
- `dotnet6` - .NET 6

## Generated Resources

1. **Lambda Function** (`aws_lambda_function`)
   - Function code placeholder
   - Environment variables
   - CloudWatch Logs integration

2. **Lambda IAM Role** (`aws_iam_role`)
   - Execution role for Lambda
   - CloudWatch Logs permissions

3. **API Gateway** (`aws_apigatewayv2_api`)
   - HTTP API
   - Default stage with auto-deploy

4. **API Gateway Integration** (`aws_apigatewayv2_integration`)
   - Lambda proxy integration
   - Request/response passthrough

5. **API Gateway Route** (`aws_apigatewayv2_route`)
   - Catch-all route ($default)
   - Routes all requests to Lambda

6. **Lambda Permission** (`aws_lambda_permission`)
   - Allows API Gateway to invoke Lambda

7. **WAF Web ACL** (`aws_wafv2_web_acl`)
   - Rate limiting rule
   - CloudWatch metrics

## Outputs

- `{name}_api_url` - API Gateway endpoint URL
- `{name}_lambda_arn` - Lambda function ARN

## Deployment

After generating Terraform files:

1. Package your Lambda code:
   ```bash
   cd infra
   zip lambda_placeholder.zip -r ../src/*
   ```

2. Initialize Terraform:
   ```bash
   terraform init
   ```

3. Preview changes:
   ```bash
   terraform plan
   ```

4. Deploy:
   ```bash
   terraform apply
   ```

## Cost Estimate

Based on moderate usage (AWS us-east-1):

- **Lambda**: $0.20 per 1M requests + $0.0000166667 per GB-second
  - Example: 1M requests/month at 512MB, 1s avg = ~$8.33/month
- **API Gateway**: $1.00 per million requests
  - Example: 1M requests/month = $1.00/month
- **WAF**: $5.00/month + $1.00 per million requests
- **CloudWatch Logs**: ~$0.50/GB ingested

**Estimated Total**: $15-20/month for 1M requests

## Security Features

- ✅ WAF with rate limiting enabled
- ✅ HTTPS only (enforced by API Gateway)
- ✅ Least-privilege IAM roles
- ✅ CloudWatch logging enabled
- ✅ VPC support (add via custom configuration)

## Example Lambda Code

### Node.js
```javascript
exports.handler = async (event) => {
    return {
        statusCode: 200,
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ message: 'Hello from Lambda!' })
    };
};
```

### Python
```python
def handler(event, context):
    return {
        'statusCode': 200,
        'headers': { 'Content-Type': 'application/json' },
        'body': '{"message": "Hello from Lambda!"}'
    }
```

## Advanced Configuration

### Add VPC Configuration

Edit generated `main.tf` to add VPC config:

```hcl
resource "aws_lambda_function" "web_api" {
  # ... existing config ...

  vpc_config {
    subnet_ids         = [aws_subnet.private.id]
    security_group_ids = [aws_security_group.lambda.id]
  }
}
```

### Add Environment Variables

Edit the Lambda resource:

```hcl
environment {
  variables = {
    ENVIRONMENT = var.environment
    DB_HOST     = aws_db_instance.main.endpoint
    API_KEY     = var.api_key
  }
}
```

### Add Custom Domain

Add to Terraform:

```hcl
resource "aws_apigatewayv2_domain_name" "api" {
  domain_name = "api.example.com"

  domain_name_configuration {
    certificate_arn = aws_acm_certificate.api.arn
    endpoint_type   = "REGIONAL"
    security_policy = "TLS_1_2"
  }
}

resource "aws_apigatewayv2_api_mapping" "api" {
  api_id      = aws_apigatewayv2_api.web_api.id
  domain_name = aws_apigatewayv2_domain_name.api.id
  stage       = aws_apigatewayv2_stage.web_api.id
}
```

## Monitoring

### CloudWatch Metrics
- Invocations
- Errors
- Duration
- Throttles
- Concurrent Executions

### CloudWatch Logs
- Function execution logs
- API Gateway access logs
- WAF logs

### Alarms

Add CloudWatch alarms:

```hcl
resource "aws_cloudwatch_metric_alarm" "lambda_errors" {
  alarm_name          = "${var.project_name}-lambda-errors"
  comparison_operator = "GreaterThanThreshold"
  evaluation_periods  = "1"
  metric_name         = "Errors"
  namespace           = "AWS/Lambda"
  period              = "60"
  statistic           = "Sum"
  threshold           = "10"
  alarm_description   = "Lambda function error rate"

  dimensions = {
    FunctionName = aws_lambda_function.web_api.function_name
  }
}
```

## Troubleshooting

### Lambda Function Not Responding
- Check CloudWatch Logs for errors
- Verify IAM role permissions
- Check timeout and memory settings

### API Gateway 403 Errors
- Verify WAF rules aren't blocking legitimate traffic
- Check API Gateway resource policies

### High Costs
- Review CloudWatch Logs retention
- Optimize Lambda memory and timeout
- Enable API Gateway caching

## References

- [AWS Lambda Documentation](https://docs.aws.amazon.com/lambda/)
- [API Gateway Documentation](https://docs.aws.amazon.com/apigateway/)
- [AWS WAF Documentation](https://docs.aws.amazon.com/waf/)