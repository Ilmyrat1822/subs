package internal

import (
	"github.com/Ilmyrat1822/subs/cmd"
	subsRouter "github.com/Ilmyrat1822/subs/internal/modules/subscription/http"
)

func InitRouters(server *cmd.Server) {
	subsRouter.InitSubscriptionRouter(server)
}
