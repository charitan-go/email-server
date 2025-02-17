package rotation

import "go.uber.org/fx"

var RotationModule = fx.Module("rotation",
	fx.Provide(
		NewRotationServer,
	),
)
