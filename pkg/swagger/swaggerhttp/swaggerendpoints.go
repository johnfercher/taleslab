package swaggerhttp

import (
	"github.com/gorilla/mux"
	"github.com/robertbakker/swaggerui"
	"go.uber.org/fx"
	"net/http"
)

var FXHandlers = fx.Options(
	fx.Invoke(DefineSwaggerEndpoint),
)

func DefineSwaggerEndpoint(router *mux.Router) {
	router.PathPrefix("/swagger/").Handler(http.StripPrefix("/swagger/", swaggerui.SwaggerFileHandler("taleslab.json")))
	router.Handle("/", http.HandlerFunc(redirectToSwagger))
	http.Handle("/", router)
}

func redirectToSwagger(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/swagger/", 301)
}
