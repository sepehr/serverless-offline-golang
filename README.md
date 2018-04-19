# Serverless Offline Golang Development Setup

A preliminary setup & skeleton to kickstart with local AWS Lambda development in Go.

It's a mixture of _Serverless_ and _SAM Local_. Using the former as deployment framework due to its wider range of
features, bigger community and its language/provider agnostic nature. And the latter only for offline Lambda + API
Gateway simulation using a docker container.

## Requirements
- Working Golang environment, preferably with `dep`
- Working Node environment to install dependencies
- Working Docker machine to invoke lambdas locally

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

## Making Changes
- Renaming lambdas requires you to update the names in `Makefile`, `serverless.yml` and `template.yml`.
- Updating APIGW endpoints requires you to update both `serverless.yml` and `template.yml` (if you want it offline).

## Additional Notes
- Included `serverless.yml` definition is originated from the official `aws-go-dep` template provided by the serverless
framework with just a minimal update to add a APIGW endpoint.
- One of the template sample lambdas has been replaced by a more useful one that can work with APIGW.
- Sample lambdas have been organized into `cmd/` directory as per [common practice](https://github.com/golang-standards/project-layout).
