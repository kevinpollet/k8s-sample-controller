#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

GOPATH=$(go env GOPATH)
CONTROLLER_TOOLS=$(grep 'sigs.k8s.io/controller-tools' go.sum | awk '{print $2}' | head -1)
CONTROLLER_TOOLS_PKG=${GOPATH}/pkg/mod/sigs.k8s.io/controller-tools@${CONTROLLER_TOOLS}

echo ">>> Using controller-tools: ${CONTROLLER_TOOLS_PKG}"

chmod +x "${CONTROLLER_TOOLS_PKG}"/.run-controller-gen.sh
chmod +x "${CONTROLLER_TOOLS_PKG}"/.run-in.sh

"${CONTROLLER_TOOLS_PKG}"/.run-controller-gen.sh crd \
    paths=./pkg/apis/... \
    output:dir=./deploy/crds
