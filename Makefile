# Makefile
# This file is used to be more language agnostic for build/automated task commands

# HELP
.DEFAULT_GOAL := help
.PHONY: help
help:
	@echo "\n\n------------------------------------------\Make Command Help\n\n"
	@echo "\nTARGETS:\n"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
	@echo ""



start: ## start sst
	npm run start -- --stage local

build: ## build sst
	npm run build -- --stage dev

deploy: ## deplooy sst to cloud
	npm run deploy -- --stage dev

remove-local: ## remove sst from cloud
	npm run remove -- --stage local

remove-dev: ## remove sst from cloud
	npm run remove -- --stage dev

remove: ## remove all env from cloud
	make remove-local;make remove-dev

test: ## test sst
	npm run test
