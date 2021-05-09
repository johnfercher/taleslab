package server

import (
	"context"
	"github.com/gorilla/mux"
	"github.com/johnfercher/taleslab/pkg/swagger/swaggerhttp"
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabhttp"
	"go.uber.org/fx"
	"net/http"
	"os"
)

var Module = fx.Options(
	fx.Invoke(
		InitServer,
	),
	taleslabhttp.FXHandlers,
	swaggerhttp.FXHandlers,
)

func InitServer(router *mux.Router, lifecycle fx.Lifecycle) {
	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}

	lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go http.ListenAndServe(":"+port, router)
			return nil
		},
		OnStop: func(ctx context.Context) error {
			panic("Close")
		},
	})
}
