#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

GOPATH=$(go env GOPATH)
CODEGEN_VERSION=$(grep 'k8s.io/code-generator' go.sum | awk '{print $2}' | head -1)
CODEGEN_PKG=${GOPATH}/pkg/mod/k8s.io/code-generator@${CODEGEN_VERSION}

echo ">>> Using code-generator: ${CODEGEN_PKG}"

chmod +x "${CODEGEN_PKG}"/generate-groups.sh

rm -rf "./pkg/apis/pk8s/v1alpha1/zz_generated.deepcopy.go"
rm -rf "./pkg/gen/client"

"${CODEGEN_PKG}"/generate-groups.sh all \
  "./pkg/gen/client" \
  "./pkg/apis" pk8s:v1alpha1 \
  --go-header-file "./hack/boilerplate.go.txt" \
  --output-base "."