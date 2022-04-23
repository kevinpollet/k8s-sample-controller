package controller

import (
	"context"

	"sigs.k8s.io/controller-runtime/pkg/builder"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

// Make sure that Reconcilier implements reconcile.Reconciler so the controller can reconcile objects.
var _ reconcile.Reconciler = &Reconciler{}

type Reconciler struct {
	Client client.Client
}

// New creates and a returns a new controller instance.
func New(mgr manager.Manager) (controller.Controller, error) {
	r := &Reconciler{
		Client: mgr.GetClient(),
	}

	return builder.ControllerManagedBy(mgr).
		Build(r)
}

func (r *Reconciler) Reconcile(_ context.Context, _ reconcile.Request) (reconcile.Result, error) {
	return reconcile.Result{}, nil
}
