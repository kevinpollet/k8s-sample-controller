package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/kevinpollet/pk8s/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/client/config"
	"sigs.k8s.io/controller-runtime/pkg/manager"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	mgr, err := manager.New(config.GetConfigOrDie(), manager.Options{})
	if err != nil {
		log.Fatalf("Unable to setup manager: %s", err)
	}

	if _, err := controller.New(mgr); err != nil {
		log.Fatalf("Unable to setup controller: %s", err)
	}

	if err := mgr.Start(ctx); err != nil {
		log.Fatalf("Unable to start manager: %s", err)
	}
}
