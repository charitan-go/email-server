package apigateway

import "go.uber.org/fx"

var ApiGatewayModule = fx.Module("apigateway",
	fx.Provide(
		NewApiGatewayRabbitmqProducer,
	),
)
