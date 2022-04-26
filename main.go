package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/kevinpollet/k8s-sample-controller/pkg/controller"
	"github.com/kevinpollet/k8s-sample-controller/pkg/k8s"
	"github.com/urfave/cli/v2"
	"sigs.k8s.io/controller-runtime/pkg/client/config"
	"sigs.k8s.io/controller-runtime/pkg/manager"
)

func main() {
	app := &cli.App{
		Name:   "k8s-sample-controller",
		Action: run,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "namespace",
				Aliases:     []string{"n"},
				Usage:       "Watch objects in the desired namespace.",
				DefaultText: "all namespaces",
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func run(cliCtx *cli.Context) error {
	ctx, cancel := signal.NotifyContext(cliCtx.Context, syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	scheme, err := k8s.NewScheme()
	if err != nil {
		return fmt.Errorf("create scheme: %w", err)
	}

	mgr, err := manager.New(config.GetConfigOrDie(), manager.Options{
		Scheme:    scheme,
		Namespace: cliCtx.String("namespace"),
	})
	if err != nil {
		return fmt.Errorf("new manager %w", err)
	}

	if _, err := controller.New(mgr); err != nil {
		return fmt.Errorf("new controller: %w", err)
	}

	return mgr.Start(ctx)
}
