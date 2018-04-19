# Serverless Offline Go Setup

A preliminary setup to kickstart with local AWS Lambda development in Go.

It's a mixture of _Serverless_ and _SAM Local_. Using the former as deployment framework due to its wider range of
features, bigger community and its language/provider agnostic nature. And the latter only for offline Lambda + API
Gateway simulation using a docker container.

## Requirements
- Working Golang environment, preferably with `dep`
- Working Node environment
- Working Docker machine

## Installation & Usage
```bash
# Install npm dependencies
npm install -g serverless aws-sam-local

# Clone the skeleton
git clone https://github.com/sepehr/serverless-offline-go.git golambda

# Compile the sample lambdas
cd golambda/
make

# Run the "apigw" lambda locally at localhost:3000
sam local start-api

# Deploy to AWS
# Requires authenticated aws-cli setup in place
sls deploy -s dev
```

### Notes
Included `serverless.yml` definition is originated from the official `aws-go-dep` template provided by the serverless.js
framework with just a minimal update to add a APIGW endpoint. Original `hello` and `world` sample lambdas have been
replaced by more useful samples.
