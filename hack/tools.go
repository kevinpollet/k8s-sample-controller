//go:build tools
// +build tools

// This package imports things required by build scripts, to force `go mod` to see them as dependencies.
// See: https://github.com/golang/go/wiki/Modules#how-can-i-track-tool-dependencies-for-a-module
package hack

import (
	_ "k8s.io/code-generator"
	_ "sigs.k8s.io/controller-tools/cmd/controller-gen"
)
