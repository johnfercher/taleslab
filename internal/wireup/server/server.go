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

	lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go listen(router)
			return nil
		},
		OnStop: func(ctx context.Context) error {
			panic("Close")
		},
	})
}

func listen(router *mux.Router) {
	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}

	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		panic(err.Error())
	}
}
