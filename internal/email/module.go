package email

import (
	"github.com/charitan-go/email-server/internal/email/service"
	"go.uber.org/fx"
)

var EmailModule = fx.Module("email",
	fx.Provide(
		service.NewEmailService,
	),
)
