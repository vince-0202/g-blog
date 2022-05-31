.PHONY: build
build-client:
	@go build -o ./bin/ github.com/vince-0202/g-blog/client/

.PHONY: test
test:
	@go test ./...