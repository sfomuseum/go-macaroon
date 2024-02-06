GOMOD=$(shell test -f "go.work" && echo "readonly" || echo "vendor")

LDFLAGS=-s -w

# https://github.com/awslabs/aws-lambda-web-adapter?tab=readme-ov-file#configurations

lambda-server:
	if test -f bootstrap; then rm -f bootstrap; fi
	if test -f macaroon-server.zip; then rm -f macaroon-server.zip; fi
	GOARCH=arm64 GOOS=linux go build -mod $(GOMOD) -ldflags="$(LDFLAGS)" -tags lambda.norpc -o bootstrap cmd/server/main.go
	zip macaroon-server.zip bootstrap
	rm -f bootstrap

