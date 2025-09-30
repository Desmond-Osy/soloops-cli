# Budget Alert Blueprint

This blueprint creates AWS Budget alerts to monitor and control infrastructure costs.

## Overview

Budget alerts help you:
- Track spending against monthly budgets
- Get notified when costs exceed thresholds
- Prevent unexpected cloud bills
- Monitor cost trends over time

## Architecture

```
AWS Cost Explorer → AWS Budgets → SNS Topic → Email Notifications
```

## Configuration

Budget is automatically configured from the `budget_usd` field in your environment:

```yaml
environments:
  - name: prod
    region: us-east-1
    budget_usd: 150    # Monthly budget in USD
    blueprints:
      # ... your blueprints
```

## Generated Resources

1. **AWS Budget** (`aws_budgets_budget`)
   - Monthly cost budget
   - Filtered by project tags
   - Multiple notification thresholds

## Notification Thresholds

By default, you'll receive alerts at:

1. **80% of budget** - Warning notification
2. **100% of budget** - Critical notification

### Example

If your budget is $150/month:
- Alert at $120 (80%)
- Alert at $150 (100%)

## Email Notifications

To receive email notifications, you need to:

1. Add your email to the budget resource in `budget.tf`:

```hcl
resource "aws_budgets_budget" "monthly" {
  # ... existing config ...

  notification {
    comparison_operator        = "GREATER_THAN"
    threshold                  = 80
    threshold_type             = "PERCENTAGE"
    notification_type          = "ACTUAL"
    subscriber_email_addresses = ["your-email@example.com"]
  }

  notification {
    comparison_operator        = "GREATER_THAN"
    threshold                  = 100
    threshold_type             = "PERCENTAGE"
    notification_type          = "ACTUAL"
    subscriber_email_addresses = ["your-email@example.com"]
  }
}
```

2. Apply the changes:

```bash
cd infra
terraform apply
```

3. Confirm the subscription:
   - Check your email for AWS Budget subscription confirmation
   - Click the confirmation link

## Advanced Configuration

### Add More Thresholds

```hcl
notification {
  comparison_operator        = "GREATER_THAN"
  threshold                  = 50
  threshold_type             = "PERCENTAGE"
  notification_type          = "ACTUAL"
  subscriber_email_addresses = ["your-email@example.com"]
}
```

### Forecast Alerts

Get notified when forecasted spending exceeds budget:

```hcl
notification {
  comparison_operator        = "GREATER_THAN"
  threshold                  = 100
  threshold_type             = "PERCENTAGE"
  notification_type          = "FORECASTED"
  subscriber_email_addresses = ["your-email@example.com"]
}
```

### Add SNS Topic for Advanced Notifications

```hcl
resource "aws_sns_topic" "budget_alerts" {
  name = "${var.project_name}-budget-alerts"
}

resource "aws_sns_topic_subscription" "budget_email" {
  topic_arn = aws_sns_topic.budget_alerts.arn
  protocol  = "email"
  endpoint  = "your-email@example.com"
}

resource "aws_sns_topic_subscription" "budget_sms" {
  topic_arn = aws_sns_topic.budget_alerts.arn
  protocol  = "sms"
  endpoint  = "+1234567890"
}

resource "aws_budgets_budget" "monthly" {
  # ... existing config ...

  notification {
    comparison_operator = "GREATER_THAN"
    threshold          = 100
    threshold_type     = "PERCENTAGE"
    notification_type  = "ACTUAL"
    subscriber_sns_topic_arns = [aws_sns_topic.budget_alerts.arn]
  }
}
```

### Multi-Environment Budgets

Each environment gets its own budget based on the `budget_usd` value:

```yaml
environments:
  - name: dev
    budget_usd: 50
    # ...
  - name: staging
    budget_usd: 100
    # ...
  - name: prod
    budget_usd: 500
    # ...
```

## Cost Filtering

Budgets automatically filter costs by:
- Project tag (matches your project name)
- Environment tag (matches environment name)

This ensures you only track costs for resources managed by SoloOps.

### Custom Cost Filters

Add additional filters:

```hcl
resource "aws_budgets_budget" "monthly" {
  # ... existing config ...

  cost_filter {
    name = "Service"
    values = [
      "Amazon Elastic Compute Cloud - Compute",
      "Amazon Simple Storage Service",
    ]
  }

  cost_filter {
    name = "LinkedAccount"
    values = ["123456789012"]
  }
}
```

## Monitoring

### View Budget Status

```bash
# AWS CLI
aws budgets describe-budgets --account-id YOUR_ACCOUNT_ID

# AWS Console
# https://console.aws.amazon.com/billing/home#/budgets
```

### CloudWatch Integration

Create CloudWatch alarms based on budget notifications:

```hcl
resource "aws_cloudwatch_metric_alarm" "budget_exceeded" {
  alarm_name          = "${var.project_name}-budget-exceeded"
  comparison_operator = "GreaterThanThreshold"
  evaluation_periods  = "1"
  metric_name         = "EstimatedCharges"
  namespace           = "AWS/Billing"
  period              = "21600"
  statistic           = "Maximum"
  threshold           = var.budget_usd
  alarm_description   = "Budget exceeded alert"
  alarm_actions       = [aws_sns_topic.budget_alerts.arn]

  dimensions = {
    Currency = "USD"
  }
}
```

## Best Practices

1. **Set Realistic Budgets**
   - Start with actual usage plus 20% buffer
   - Adjust monthly based on trends

2. **Use Multiple Thresholds**
   - 50% - Early warning
   - 80% - Action required
   - 100% - Critical alert

3. **Tag All Resources**
   - Ensure resources have Project and Environment tags
   - Use consistent tagging across all services

4. **Review Regularly**
   - Check budget reports weekly
   - Analyze cost trends monthly
   - Adjust budgets quarterly

5. **Automate Responses**
   - Use Lambda to respond to alerts
   - Implement auto-scaling limits
   - Set up resource quotas

## Cost Optimization Tips

If you're approaching your budget:

1. **Review CloudWatch Logs**
   - Reduce retention periods
   - Use S3 for long-term log storage

2. **Optimize Lambda**
   - Right-size memory allocation
   - Reduce timeout values
   - Use provisioned concurrency sparingly

3. **CloudFront Caching**
   - Increase cache TTLs
   - Enable compression
   - Use CloudFront Functions instead of Lambda@Edge

4. **S3 Storage**
   - Use appropriate storage class
   - Enable lifecycle policies
   - Delete unused objects

5. **API Gateway**
   - Enable caching
   - Use HTTP APIs instead of REST APIs
   - Optimize request/response sizes

## Troubleshooting

### Not Receiving Notifications

- Check email subscription is confirmed
- Verify email address in budget configuration
- Check spam folder
- Ensure SNS topic has correct permissions

### Budget Not Tracking Costs

- Verify resource tags are correct
- Check cost filter configuration
- Wait 24 hours for costs to appear
- Ensure billing is enabled in AWS account

### Incorrect Cost Amounts

- Budgets show estimated costs
- Actual costs are finalized after month-end
- Some services have delayed cost reporting
- Credits and refunds may not appear immediately

## Additional Resources

- [AWS Budgets Documentation](https://docs.aws.amazon.com/cost-management/latest/userguide/budgets-managing-costs.html)
- [AWS Cost Optimization](https://aws.amazon.com/aws-cost-management/aws-cost-optimization/)
- [AWS Pricing Calculator](https://calculator.aws/)

## Example Alert Email

When a threshold is exceeded, you'll receive an email like:

```
Subject: AWS Budget Alert: my-project-prod-monthly has exceeded 80% of your budget

Your AWS Budget my-project-prod-monthly has exceeded 80% of your $150.00 budget.

Current spending: $120.50
Budget: $150.00
Percentage: 80.3%

Period: March 2025

View details: [Link to AWS Console]
```