package main

import (
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"github.com/johnfercher/taleslab/internal/api/apiencodes"
	"github.com/johnfercher/taleslab/pkg/api/forest/foresthttp"
	"github.com/johnfercher/taleslab/pkg/api/forest/forestservices"
	"github.com/johnfercher/taleslab/pkg/slabcompressor"
	"github.com/johnfercher/taleslab/pkg/slabdecoder"
	"net/http"
)

func main() {

	encoder := slabdecoder.NewEncoder(slabcompressor.New())

	forestService := forestservices.NewForestService(encoder)

	serverOptions := []httptransport.ServerOption{
		httptransport.ServerErrorEncoder(apiencodes.EncodeError),
	}

	generateForestEndpoint := httptransport.NewServer(apiencodes.LogRequest(foresthttp.MakeGenerateForest(forestService)),
		foresthttp.DecodeForestRequest,
		apiencodes.EncodeResponse,
		serverOptions...,
	)

	router := mux.NewRouter()
	router.Handle("/api/generate/forest", generateForestEndpoint)

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		print(err.Error())
		return
	}

}
