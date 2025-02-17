package key

import (
	"github.com/charitan-go/key-server/internal/key/service"
	"go.uber.org/fx"
)

var KeyModule = fx.Module("key",
	fx.Provide(
		service.NewKeyService,
	),
)
