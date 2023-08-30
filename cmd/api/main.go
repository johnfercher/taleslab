// Package classification TaleSlab
//
// # This is an API designed to generate TaleSpire slabs dinamically
//
// Terms Of Service:
//
//	Schemes: https, http
//	Host: taleslab.herokuapp.com
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
	"github.com/johnfercher/taleslab/internal/helper/bytecompressor"
	"github.com/johnfercher/taleslab/internal/helper/talespireadapter/talespirecoder"
	"github.com/johnfercher/taleslab/internal/httprouter"
	"github.com/johnfercher/taleslab/internal/wireup/server"
	"github.com/johnfercher/taleslab/pkg/taleslab"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		taleslab.Module,
		bytecompressor.Module,
		talespirecoder.Module,
		httprouter.Module,
		server.Module,
	).Run()
}
