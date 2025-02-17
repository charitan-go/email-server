package rabbitmq

import "go.uber.org/fx"

var RabbitmqModule = fx.Module("rabbitmq",
	fx.Provide(
		NewRabbitmqService,
	),
)
