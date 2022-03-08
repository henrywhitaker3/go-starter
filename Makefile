compile:
	CGO_ENABLED=0 GOOS=$(os) GOARCH=$(arch) go build -o bin/hermes-$(version)-$(os)-$(arch) main.go

test:
	go test ./...