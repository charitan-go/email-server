package inbucket

import "go.uber.org/fx"

var InbucketModule = fx.Module("inbucket",
	fx.Provide(
		NewInbucketService,
	),
)
