#
# Makefile fragment for Generate
#

GO           ?= go

PACKAGES ?= $(shell $(GO) list ./...)

GOTOOLS += github.com/newrelic/tutone/cmd/tutone

# Generate then lint fixes
generate: generate-run generate-schema-fetch generate-tutone lint-fix

generate-schema-fetch:
	@echo "=== $(PROJECT_NAME) === [ schema-fetch     ]: Running tutone fetch..."
	tutone -c .tutone.yml fetch -l debug

generate-tutone:
	@echo "=== $(PROJECT_NAME) === [ generate-tutone  ]: Running tutone generate..."
	@for p in $(PACKAGES); do \
		echo "=== $(PROJECT_NAME) === [ generate         ]:     $$p"; \
			tutone -c .tutone.yml generate -d .generate.yml -l debug -p $$p; \
	done

generate-run: tools-compile generate-tutone
	@echo "=== $(PROJECT_NAME) === [ generate         ]: Running generate..."
	@for p in $(PACKAGES); do \
		echo "=== $(PROJECT_NAME) === [ generate         ]:     $$p"; \
			$(GO) generate -x $$p ; \
	done

.PHONY: generate generate-run
