package httprouter

import (
	"github.com/gorilla/mux"
	"go.uber.org/fx"
)

func New() *mux.Router {
	router := mux.NewRouter()
	return router
}

var Module = fx.Option(
	fx.Provide(New),
)
