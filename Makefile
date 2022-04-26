.PHONY: help
help:  ## Displays this help.
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_0-9-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

.PHONY: image
image: ## Builds the Docker image.
	docker build . -t kevinpollet/k8s-sample-controller

.PHONY: update-codegen
update-codegen: ## Generates Kubernetes client and deepcopy functions.
	./scripts/update-codegen.sh

.PHONY: verify-codegen
verify-codegen: ## Verifies that generated Kubernetes client and deepcopy functions.
	./scripts/verify-codegen.sh

.PHONY: update-crds
update-crds: ## Generates CRDs (Custom Resource Definitions).
	./scripts/update-crds.sh

.PHONY: verify-crds
verify-crds: ## Verifies that generated CRDs are up-to-date.
	./scripts/verify-crds.sh