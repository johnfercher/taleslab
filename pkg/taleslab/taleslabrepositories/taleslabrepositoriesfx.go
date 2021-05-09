package taleslabrepositories

import "go.uber.org/fx"

var Module = fx.Option(
	fx.Provide(NewBiomeRepository, NewPropRepository),
)
