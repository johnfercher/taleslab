//	Package classification TaleSlab
//
//	This is an API designed to generate TaleSpire slabs dinamically
//
//
//	Terms Of Service:
//
//		Schemes: https, http
//		Host: taleslab.herokuapp.com
//		Version: 1.0.0
//
//		Consumes:
//		- application/json
//
//		Produces:
//		- application/json
//
//swagger:meta
//go:generate swagger generate spec -o ../taleslab.json
package main

import (
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"github.com/johnfercher/taleslab/internal/api/apiencodes"
	"github.com/johnfercher/taleslab/internal/bytecompressor"
	"github.com/johnfercher/taleslab/pkg/assetloader"
	"github.com/johnfercher/taleslab/pkg/biomeloader"
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslab/taleslabhttp"
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslab/taleslabservices"
	"github.com/johnfercher/taleslab/pkg/talespire/talespirecoder"
	"github.com/robertbakker/swaggerui"
	"net/http"
	"os"
)

func main() {
	byteCompressor := bytecompressor.New()
	encoder := talespirecoder.NewEncoder(byteCompressor)
	assetLoader := assetloader.NewAssetLoader()
	biomeLoader := biomeloader.NewBiomeLoader(assetLoader)
	mapService := taleslabservices.NewMapService(biomeLoader, encoder)

	serverOptions := []httptransport.ServerOption{
		httptransport.ServerErrorEncoder(apiencodes.EncodeError),
	}

	generateMapEndpoint := httptransport.NewServer(apiencodes.LogRequest(taleslabhttp.MakeGenerateMap(mapService)),
		taleslabhttp.DecodeMapRequest,
		apiencodes.EncodeResponse,
		serverOptions...,
	)

	router := mux.NewRouter()
	// swagger:operation POST /api/generate/map map
	// ---
	// summary: Generate a new map, based on the input parameters
	// description: The biome you select will change the ground tile and tree type.
	// produces:
	// - application/json
	// parameters:
	// - name: body
	//   description: Input parameters for map generation
	//   in: body
	//   required: true
	//   schema:
	//     "$ref": "#/definitions/Map"
	// responses:
	//   "200":
	//     "$ref": "#/responses/mapRes"
	//   "400":
	//     "$ref": "#/responses/errRes"
	//   "404":
	//     "$ref": "#/responses/errRes"
	router.Handle("/api/generate/map", generateMapEndpoint)

	router.PathPrefix("/swagger/").Handler(http.StripPrefix("/swagger/", swaggerui.SwaggerFileHandler("taleslab.json")))

	router.Handle("/", http.HandlerFunc(redirectToSwagger))

	http.Handle("/", router)

	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}

	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		print(err.Error())
		return
	}

}

func redirectToSwagger(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/swagger/", 301)
}
