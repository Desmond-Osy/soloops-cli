/**
 * Example Lambda Handler - Node.js
 *
 * This is a sample Lambda function that demonstrates:
 * - Handling HTTP events from API Gateway
 * - Environment variable access
 * - Error handling
 * - JSON response formatting
 */

exports.handler = async (event, context) => {
    console.log('Event:', JSON.stringify(event, null, 2));
    console.log('Environment:', process.env.ENVIRONMENT);

    try {
        // Parse request
        const method = event.requestContext.http.method;
        const path = event.requestContext.http.path;

        // Route handling
        if (method === 'GET' && path === '/health') {
            return healthCheck();
        }

        if (method === 'GET' && path === '/') {
            return welcomeMessage();
        }

        if (method === 'POST' && path === '/api/data') {
            const body = JSON.parse(event.body || '{}');
            return processData(body);
        }

        // 404 for unknown routes
        return {
            statusCode: 404,
            headers: {
                'Content-Type': 'application/json',
                'Access-Control-Allow-Origin': '*'
            },
            body: JSON.stringify({
                error: 'Not Found',
                message: `Route ${method} ${path} not found`
            })
        };

    } catch (error) {
        console.error('Error:', error);

        return {
            statusCode: 500,
            headers: {
                'Content-Type': 'application/json',
                'Access-Control-Allow-Origin': '*'
            },
            body: JSON.stringify({
                error: 'Internal Server Error',
                message: error.message
            })
        };
    }
};

/**
 * Health check endpoint
 */
function healthCheck() {
    return {
        statusCode: 200,
        headers: {
            'Content-Type': 'application/json',
            'Access-Control-Allow-Origin': '*'
        },
        body: JSON.stringify({
            status: 'healthy',
            timestamp: new Date().toISOString(),
            environment: process.env.ENVIRONMENT || 'unknown'
        })
    };
}

/**
 * Welcome message
 */
function welcomeMessage() {
    return {
        statusCode: 200,
        headers: {
            'Content-Type': 'application/json',
            'Access-Control-Allow-Origin': '*'
        },
        body: JSON.stringify({
            message: 'Welcome to your Serverless API!',
            version: '1.0.0',
            endpoints: {
                health: 'GET /health',
                data: 'POST /api/data'
            }
        })
    };
}

/**
 * Process data example
 */
function processData(data) {
    // Validate input
    if (!data || typeof data !== 'object') {
        return {
            statusCode: 400,
            headers: {
                'Content-Type': 'application/json',
                'Access-Control-Allow-Origin': '*'
            },
            body: JSON.stringify({
                error: 'Bad Request',
                message: 'Invalid JSON data'
            })
        };
    }

    // Process data
    const result = {
        received: data,
        processed: true,
        timestamp: new Date().toISOString()
    };

    return {
        statusCode: 200,
        headers: {
            'Content-Type': 'application/json',
            'Access-Control-Allow-Origin': '*'
        },
        body: JSON.stringify(result)
    };
}