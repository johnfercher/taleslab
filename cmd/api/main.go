// Package classification TaleSlab
//
// # This is an API designed to generate TaleSpire slabs dinamically
//
// Terms Of Service:
//
//	Schemes: https, http
//	Host: localhost:5000
//	Version: 1.0.0
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//swagger:meta
//go:generate swagger generate spec -o ../../taleslab.json
package main

import (
	"github.com/johnfercher/talescoder/pkg/decoder"
	"github.com/johnfercher/talescoder/pkg/encoder"
	"github.com/johnfercher/taleslab/internal/httprouter"
	"github.com/johnfercher/taleslab/internal/wireup/server"
	"github.com/johnfercher/taleslab/pkg/taleslab"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		taleslab.Module,
		httprouter.Module,
		server.Module,
		fx.Provide(encoder.NewEncoder, decoder.NewDecoder),
	).Run()
}
