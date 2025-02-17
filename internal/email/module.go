package email

import (
	"github.com/charitan-go/email-server/internal/email/service"
	"go.uber.org/fx"
)

var KeyModule = fx.Module("email",
	fx.Provide(
		service.NewKeyService,
	),
)
