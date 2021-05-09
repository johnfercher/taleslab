package taleslab

import (
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabrepositories"
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabservices"
	"go.uber.org/fx"
)

var Module = fx.Options(
	taleslabrepositories.Module,
	taleslabservices.Module,
)
