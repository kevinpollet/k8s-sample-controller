#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

GOPATH=$(go env GOPATH)
CTOOLS_VERSION=$(grep 'sigs.k8s.io/controller-tools' go.sum | awk '{print $2}' | head -1)
CTOOLS_PKG=${GOPATH}/pkg/mod/sigs.k8s.io/controller-tools@${CTOOLS_VERSION}

echo ">>> Using controller-tools: ${CTOOLS_PKG}"

chmod +x "${CTOOLS_PKG}"/.run-controller-gen.sh
chmod +x "${CTOOLS_PKG}"/.run-in.sh

"${CTOOLS_PKG}"/.run-controller-gen.sh crd \
  paths=./pkg/apis/... \
  output:dir=./deploy/crds
