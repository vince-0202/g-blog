.PHONY: build
build:
	@go build -o ./bin/ github.com/vince-0202/g-blog/client/

.PHONY: test
test:
	@go test ./...

.PHONY: build-admin
build-admin:
	@go build -o ./bin/ github.com/vince-0202/g-blog/admin/
	