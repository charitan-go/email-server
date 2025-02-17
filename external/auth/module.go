package auth

import (
	"go.uber.org/fx"
)

var AuthModule = fx.Module("auth",
	fx.Provide(
		NewAuthRabbitmqProducer,
	),
)
