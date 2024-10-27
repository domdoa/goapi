build:
	@go build -o goapi
run: build
	@./goapi
test:
	@go test -v ./...