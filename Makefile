GOMOD=$(shell test -f "go.work" && echo "readonly" || echo "vendor")

LDFLAGS=-s -w

cli:
	go build -mod $(GOMOD) -ldflags="$(LDFLAGS)" -o bin/new-macaroon cmd/new-macaroon/main.go
	go build -mod $(GOMOD) -ldflags="$(LDFLAGS)" -o bin/tp-ensure-account cmd/tp-ensure-account/main.go
	go build -mod $(GOMOD) -ldflags="$(LDFLAGS)" -o bin/signing-key cmd/signing-key/main.go
	go build -mod $(GOMOD) -ldflags="$(LDFLAGS)" -o bin/encryption-key cmd/encryption-key/main.go
