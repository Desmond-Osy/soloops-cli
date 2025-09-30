#!/bin/bash

# Static Site Deployment Script
# This script uploads your website to S3 and invalidates CloudFront cache

set -e

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# Check if AWS CLI is installed
if ! command -v aws &> /dev/null; then
    echo -e "${RED}Error: AWS CLI is not installed${NC}"
    echo "Install it from: https://aws.amazon.com/cli/"
    exit 1
fi

# Get bucket name and distribution ID from Terraform outputs
echo -e "${YELLOW}Getting infrastructure details...${NC}"

cd infra

BUCKET_NAME=$(terraform output -raw static_site_bucket_name 2>/dev/null || echo "")
DISTRIBUTION_ID=$(terraform output -raw static_site_cloudfront_distribution_id 2>/dev/null || echo "")

if [ -z "$BUCKET_NAME" ]; then
    echo -e "${RED}Error: Could not get S3 bucket name from Terraform outputs${NC}"
    echo "Make sure you've run 'terraform apply' first"
    exit 1
fi

echo -e "${GREEN}✓ Bucket: $BUCKET_NAME${NC}"

cd ..

# Upload files to S3
echo -e "${YELLOW}Uploading files to S3...${NC}"

aws s3 sync ./example-website s3://$BUCKET_NAME/ \
    --delete \
    --cache-control "public, max-age=31536000" \
    --exclude "*.html" \
    --exclude "*.json"

# Upload HTML files with shorter cache
aws s3 sync ./example-website s3://$BUCKET_NAME/ \
    --exclude "*" \
    --include "*.html" \
    --cache-control "public, max-age=300"

echo -e "${GREEN}✓ Files uploaded successfully${NC}"

# Invalidate CloudFront cache
if [ -n "$DISTRIBUTION_ID" ]; then
    echo -e "${YELLOW}Invalidating CloudFront cache...${NC}"

    INVALIDATION_ID=$(aws cloudfront create-invalidation \
        --distribution-id $DISTRIBUTION_ID \
        --paths "/*" \
        --query 'Invalidation.Id' \
        --output text)

    echo -e "${GREEN}✓ Invalidation created: $INVALIDATION_ID${NC}"
    echo -e "${YELLOW}Waiting for invalidation to complete (this may take a few minutes)...${NC}"

    aws cloudfront wait invalidation-completed \
        --distribution-id $DISTRIBUTION_ID \
        --id $INVALIDATION_ID

    echo -e "${GREEN}✓ Invalidation completed${NC}"
else
    echo -e "${YELLOW}⚠ CloudFront distribution ID not found, skipping cache invalidation${NC}"
fi

# Get CloudFront URL
CLOUDFRONT_URL=$(cd infra && terraform output -raw static_site_cloudfront_url 2>/dev/null || echo "")

if [ -n "$CLOUDFRONT_URL" ]; then
    echo ""
    echo -e "${GREEN}━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━${NC}"
    echo -e "${GREEN}✓ Deployment complete!${NC}"
    echo -e "${GREEN}━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━${NC}"
    echo ""
    echo -e "Your site is live at: ${GREEN}https://$CLOUDFRONT_URL${NC}"
    echo ""
else
    echo -e "${GREEN}✓ Deployment complete!${NC}"
fi