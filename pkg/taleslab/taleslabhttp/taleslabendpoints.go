package taleslabhttp

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"github.com/johnfercher/taleslab/internal/api/apiencodes"
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabdomain/taleslabservices"
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabdto"
	"go.uber.org/fx"
)

var FXHandlers = fx.Options(
	fx.Invoke(DefineGetGenerationsCountEndpoint, DefineGenerateMapEndpoint),
)

var serverOptions = []httptransport.ServerOption{
	httptransport.ServerErrorEncoder(apiencodes.EncodeError),
}

func MakeGenerateMap(service taleslabservices.SlabGenerator) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		inputMap := request.(*taleslabdto.MapDtoRequest)
		return service.Generate(ctx, inputMap)
	}
}

func MakeGetGenerationsCount(service taleslabservices.SlabGenerator) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		return service.GetGenerationsCount(ctx)
	}
}

func DefineGenerateMapEndpoint(router *mux.Router, mapService taleslabservices.SlabGenerator) {
	generateMapEndpoint := httptransport.NewServer(apiencodes.LogRequest(MakeGenerateMap(mapService)),
		DecodeMapRequest,
		apiencodes.EncodeResponse,
		serverOptions...,
	)

	// swagger:operation POST /api/generate/map map_generation
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
	//     "$ref": "#/definitions/MapDtoRequest"
	// responses:
	//   "200":
	//     "$ref": "#/responses/swaggMapRes"
	//   "400":
	//     "$ref": "#/responses/errRes"
	//   "404":
	//     "$ref": "#/responses/errRes"
	router.Handle("/api/generate/map", generateMapEndpoint)
}

func DefineGetGenerationsCountEndpoint(router *mux.Router, mapService taleslabservices.SlabGenerator) {

	getGenerationsCountEndpoint := httptransport.NewServer(apiencodes.LogRequest(MakeGetGenerationsCount(mapService)),
		DecodeNothing,
		apiencodes.EncodeResponse,
		serverOptions...,
	)

	// swagger:operation GET /api/count get_count
	// ---
	// summary: Get quantity of maps generated
	// description: Get how many maps were generated
	// produces:
	// - application/json
	// responses:
	//   "200":
	//     "$ref": "#/responses/swaggCountRes"
	//   "400":
	//     "$ref": "#/responses/errRes"
	//   "404":
	//     "$ref": "#/responses/errRes"
	router.Handle("/api/count", getGenerationsCountEndpoint)

}
