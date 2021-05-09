package talespirecoder

import "go.uber.org/fx"

var Module = fx.Option(
	fx.Provide(NewEncoder, NewDecoder),
)
