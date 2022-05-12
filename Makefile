
# HELP
.DEFAULT_GOAL := help
.PHONY: help
help:
	@echo "\n\n------------------------------------------\Make Command Help\n\n"
	@echo "\nTARGETS:\n"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
	@echo ""



start: ## start sst
	npm start

build: ## build sst
	npm run build

deploy: ## deplooy sst to cloud
	npm run deploy

remove: ## remove sst from cloud
	npm run remove

test: ## test sst
	npm run test
