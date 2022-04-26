package k8s

import (
	"github.com/kevinpollet/k8s-sample-controller/pkg/apis/sample/v1alpha1"
	"k8s.io/apimachinery/pkg/runtime"
)

// NewScheme returns a scheme including the manipulated API types.
func NewScheme() (*runtime.Scheme, error) {
	scheme := runtime.NewScheme()
	schemeBuilder := runtime.SchemeBuilder{
		v1alpha1.AddToScheme,
	}

	if err := schemeBuilder.AddToScheme(scheme); err != nil {
		return nil, err
	}
	return scheme, nil
}
