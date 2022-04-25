#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

SCRIPT_ROOT=$(dirname "${BASH_SOURCE[0]}")/..
GOPATH=$(go env GOPATH)
CODEGEN_VERSION=$(go list -m -f "{{ .Version }}" "k8s.io/code-generator")
CODEGEN_PKG=${GOPATH}/pkg/mod/k8s.io/code-generator@${CODEGEN_VERSION}
_tmp=$(mktemp -d)

cleanup() {
    rm -rf ${_tmp}
}
trap "cleanup" EXIT SIGINT

cleanup

echo ">>> Using code-generator: ${CODEGEN_PKG}"

chmod +x "${CODEGEN_PKG}"/generate-groups.sh

rm -rf "${SCRIPT_ROOT}/pkg/apis/sample/v1alpha1/zz_generated.deepcopy.go"
rm -rf "${SCRIPT_ROOT}/pkg/client"

"${CODEGEN_PKG}"/generate-groups.sh all \
  "github.com/kevinpollet/k8s-sample-controller/pkg/client" \
  "github.com/kevinpollet/k8s-sample-controller/pkg/apis" \
  sample:v1alpha1 \
  --go-header-file "${SCRIPT_ROOT}/scripts/boilerplate.go.txt" \
  --output-base "${_tmp}"

cp -a "${_tmp}/github.com/kevinpollet/k8s-sample-controller"/* "${SCRIPT_ROOT}"
