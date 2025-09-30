"""
Example Lambda Handler - Python

This is a sample Lambda function that demonstrates:
- Handling HTTP events from API Gateway
- Environment variable access
- Error handling
- JSON response formatting
"""

import json
import os
from datetime import datetime


def handler(event, context):
    """Main Lambda handler function"""
    print(f"Event: {json.dumps(event)}")
    print(f"Environment: {os.environ.get('ENVIRONMENT', 'unknown')}")

    try:
        # Parse request
        http_context = event.get('requestContext', {}).get('http', {})
        method = http_context.get('method', 'GET')
        path = http_context.get('path', '/')

        # Route handling
        if method == 'GET' and path == '/health':
            return health_check()

        if method == 'GET' and path == '/':
            return welcome_message()

        if method == 'POST' and path == '/api/data':
            body = json.loads(event.get('body', '{}'))
            return process_data(body)

        # 404 for unknown routes
        return {
            'statusCode': 404,
            'headers': {
                'Content-Type': 'application/json',
                'Access-Control-Allow-Origin': '*'
            },
            'body': json.dumps({
                'error': 'Not Found',
                'message': f'Route {method} {path} not found'
            })
        }

    except Exception as e:
        print(f"Error: {str(e)}")

        return {
            'statusCode': 500,
            'headers': {
                'Content-Type': 'application/json',
                'Access-Control-Allow-Origin': '*'
            },
            'body': json.dumps({
                'error': 'Internal Server Error',
                'message': str(e)
            })
        }


def health_check():
    """Health check endpoint"""
    return {
        'statusCode': 200,
        'headers': {
            'Content-Type': 'application/json',
            'Access-Control-Allow-Origin': '*'
        },
        'body': json.dumps({
            'status': 'healthy',
            'timestamp': datetime.utcnow().isoformat(),
            'environment': os.environ.get('ENVIRONMENT', 'unknown')
        })
    }


def welcome_message():
    """Welcome message"""
    return {
        'statusCode': 200,
        'headers': {
            'Content-Type': 'application/json',
            'Access-Control-Allow-Origin': '*'
        },
        'body': json.dumps({
            'message': 'Welcome to your Serverless API!',
            'version': '1.0.0',
            'endpoints': {
                'health': 'GET /health',
                'data': 'POST /api/data'
            }
        })
    }


def process_data(data):
    """Process data example"""
    # Validate input
    if not data or not isinstance(data, dict):
        return {
            'statusCode': 400,
            'headers': {
                'Content-Type': 'application/json',
                'Access-Control-Allow-Origin': '*'
            },
            'body': json.dumps({
                'error': 'Bad Request',
                'message': 'Invalid JSON data'
            })
        }

    # Process data
    result = {
        'received': data,
        'processed': True,
        'timestamp': datetime.utcnow().isoformat()
    }

    return {
        'statusCode': 200,
        'headers': {
            'Content-Type': 'application/json',
            'Access-Control-Allow-Origin': '*'
        },
        'body': json.dumps(result)
    }