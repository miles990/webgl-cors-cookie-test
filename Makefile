# VARIABLES
GOPATH=/home/alex/go
PACKAGE=webgl-cors-cookie-test
BINARY_NAME=webgl-cors-cookie-test

default: usage

clean: ## Trash binary files
	@echo "--> cleaning..."
	@go clean || (echo "Unable to clean project" && exit 1)
	@rm -rf $(GOPATH)/bin/$(BINARY_NAME) 2> /dev/null	
	@echo "Clean OK"

test: ## Run all tests
	@echo "--> testing..."
	@go test -v $(PACKAGE)/tests/...

install: clean ## Compile sources and build binary
	@echo "--> updating repo..."
	@cd $(GOPATH)/src/$(PACKAGE)
	@git checkout -f develop
	@git reset --hard HEAD
	@git pull --all
	@echo "--> installing package..."	
	@go get -v github.com/beego/bee
	@go get -v github.com/astaxie/beego
	@go install $(PACKAGE) || (echo "Compilation error" && exit 1)
	@echo "Install OK"

run: install ## Run your application
	@echo "--> running application..."
	@$(GOPATH)/bin/$(BINARY_NAME)

usage: ## List available targets
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
