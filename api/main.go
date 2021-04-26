package main

import (
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"github.com/johnfercher/taleslab/internal/api/apiencodes"
	"github.com/johnfercher/taleslab/pkg/api/taleslab/taleslabhttp"
	"github.com/johnfercher/taleslab/pkg/api/taleslab/taleslabservices"
	"github.com/johnfercher/taleslab/pkg/slabcompressor"
	"github.com/johnfercher/taleslab/pkg/slabdecoder"
	"net/http"
	"os"
)

func main() {

	encoder := slabdecoder.NewEncoder(slabcompressor.New())

	mapservice := taleslabservices.NewMapService(encoder)

	serverOptions := []httptransport.ServerOption{
		httptransport.ServerErrorEncoder(apiencodes.EncodeError),
	}

	generateMapEndpoint := httptransport.NewServer(apiencodes.LogRequest(taleslabhttp.MakeGenerateMap(mapservice)),
		taleslabhttp.DecodeMapRequest,
		apiencodes.EncodeResponse,
		serverOptions...,
	)

	router := mux.NewRouter()
	router.Handle("/api/generate/map", generateMapEndpoint)

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
