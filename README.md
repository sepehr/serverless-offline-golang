你好！
很冒昧用这样的方式来和你沟通，如有打扰请忽略我的提交哈。我是光年实验室（gnlab.com）的HR，在招Golang开发工程师，我们是一个技术型团队，技术氛围非常好。全职和兼职都可以，不过最好是全职，工作地点杭州。
我们公司是做流量增长的，Golang负责开发SAAS平台的应用，我们做的很多应用是全新的，工作非常有挑战也很有意思，是国内很多大厂的顾问。
如果有兴趣的话加我微信：13515810775  ，也可以访问 https://gnlab.com/，联系客服转发给HR。
# Serverless Offline Golang Development Setup

A preliminary setup & skeleton to kickstart with local AWS Lambda development in Go.

It's a mixture of _Serverless_ and _SAM Local_. Using the former as deployment framework due to its wider range of
features, bigger community and its language/provider agnostic nature. And the latter only for offline Lambda + API
Gateway simulation using a docker container.

## Requirements
- Working Golang environment, preferably with `dep`
- Working Node environment to install dependencies
- Working Docker machine to invoke lambdas locally

## Installation & usage
```bash
# Install npm dependencies
npm install -g serverless aws-sam-local

# Clone the skeleton
git clone https://github.com/sepehr/serverless-offline-go.git golambda

# Compile the sample lambdas
cd golambda/
make

# Invoke the "apigw" sample lambda locally at localhost:3000
sam local start-api

# Invoke the "Vanilla" sample lambda locally using a custom event payload file
sam local invoke "Vanilla" -e path/to/event.json
# Or using an event payload from the stdin
 echo '{"message": "Hey, are you there?" }' | sam local invoke "Vanilla"

# Deploy to AWS
# Requires authenticated aws-cli setup in place
sls deploy -s dev
```

## Making changes
- Renaming lambdas requires you to update the names in `Makefile`, `serverless.yml` and `template.yml`.
- Updating APIGW endpoints requires you to update both `serverless.yml` and `template.yml` (if you want it offline).

## Additional Notes
- Included `serverless.yml` definition is originated from the official `aws-go-dep` template provided by the serverless
framework with just a minimal update to add a APIGW endpoint.
- One of the template sample lambdas has been replaced by a more useful one that can work with APIGW.
- Sample lambdas have been organized into `cmd/` directory as per [common practice](https://github.com/golang-standards/project-layout).

- - -

[![asciicast](https://asciinema.org/a/177254.png)](https://asciinema.org/a/177254)
