.PHONY: lint install-lint goimports install-goimports format clean

# 版本定义
GOLANGCI_LINT_VERSION ?= v1.61.0
GOIMPORTS_VERSION ?= v0.25.0

# Go 相关环境变量
GOBIN ?= $(shell go env GOPATH)/bin
GOIMPORTS ?= goimports
GOFMT ?= gofmt

# 检查是否安装了 golangci-lint
GOLANGCI_LINT ?= $(shell which golangci-lint || echo "$(GOBIN)/golangci-lint")

# 安装 golangci-lint（如果未安装）
install-lint:
	@if ! command -v golangci-lint >/dev/null 2>&1; then \
		echo "Installing golangci-lint $(GOLANGCI_LINT_VERSION)..."; \
		curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(GOBIN) $(GOLANGCI_LINT_VERSION); \
	else \
		echo "golangci-lint already installed"; \
	fi

# 运行 lint 检查
lint: install-lint
	@echo "Running golangci-lint..."
	@$(GOLANGCI_LINT) run ./...

# goimports 执行路径
GOIMPORTS ?= $(shell which goimports || echo "$(GOBIN)/goimports")

# 安装 goimports（如果未安装）
install-goimports:
	@if ! command -v goimports >/dev/null 2>&1; then \
		echo "Installing goimports $(GOIMPORTS_VERSION)..."; \
		go install golang.org/x/tools/cmd/goimports@$(GOIMPORTS_VERSION); \
	else \
		echo "goimports already installed"; \
	fi

# 运行 goimports 和 gofmt 格式化代码
format:
	@echo "Running goimports..."
	@$(GOIMPORTS) -w .
	@echo "Running gofmt..."
	@$(GOFMT) -w .
	@echo "Code formatting completed."
# 清理临时文件
clean:
	@rm -rf ./tmp