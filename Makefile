run:
	@go run cmd/sso/main.go --config=./config/local.yaml

build:
	@go build -o sso