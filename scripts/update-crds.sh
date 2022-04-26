#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

GOPATH=$(go env GOPATH)
CONTROLLER_TOOLS_VERSION=$(go list -m -f "{{ .Version }}" "sigs.k8s.io/controller-tools")
CONTROLLER_TOOLS_PKG=${GOPATH}/pkg/mod/sigs.k8s.io/controller-tools@${CONTROLLER_TOOLS_VERSION}

echo ">>> Using controller-tools: ${CONTROLLER_TOOLS_PKG}"

chmod +x "${CONTROLLER_TOOLS_PKG}"/.run-controller-gen.sh
chmod +x "${CONTROLLER_TOOLS_PKG}"/.run-in.sh

rm -rf "./deploy/crds"

"${CONTROLLER_TOOLS_PKG}"/.run-controller-gen.sh crd \
  paths=./pkg/apis/... \
  output:dir=./deploy/crds \
