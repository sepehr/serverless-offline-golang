build:
	dep ensure
	env GOOS=linux go build -ldflags="-s -w" -o bin/apigw cmd/apigw/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/vanilla cmd/vanilla/main.go
