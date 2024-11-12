.PHONY: build
build:
	@rm webook || true
	@GOOS=linux GOARCH=amd64 go build -tags=k8s -o webook ./cmd/
	@docker rmi -f novo/webook:v0.0.1
	@docker build -t novo/webook:v0.0.1 .
.PHONY: mock
mock:
	# window下需要单引号包裹路径
	@mockgen -source='internal/repository/user.go' -destination='internal/repository/user_mock.go' -package=repository