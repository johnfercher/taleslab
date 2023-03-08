package taleslabservices

import "go.uber.org/fx"

var Module = fx.Option(
	fx.Provide(NewMapService, NewAssetsGenerator),
)
