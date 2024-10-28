.PHONY: build
build:
	@rm webook || true
	@GOOS=linux GOARCH=amd64 go build -tags=k8s -o webook ./cmd/
	@docker rmi -f novo/webook:v0.0.1
	@docker build -t novo/webook:v0.0.1 .