package http

import (
	"github.com/Ilmyrat1822/subs/cmd"
	"github.com/Ilmyrat1822/subs/internal/modules/subscription/handler"
	"github.com/Ilmyrat1822/subs/internal/modules/subscription/repository"
	"github.com/Ilmyrat1822/subs/internal/modules/subscription/service"
)

func InitSubscriptionRouter(server *cmd.Server) {
	subsRepository := repository.NewSubscriptionRepository(server.Database)
	subsService := service.NewSubscriptionService(subsRepository)
	subsHandler := handler.NewSubscriptionHandler(subsService)

	subsRouter := server.Echo.Group("/api/subs")
	subsRouter.GET("/list", subsHandler.List)
	subsRouter.POST("", subsHandler.Create)
	subsRouter.GET("/total", subsHandler.TotalCost)
	subsRouter.GET("/:id", subsHandler.Get)
	subsRouter.PUT("/:id", subsHandler.Update)
	subsRouter.DELETE("/:id", subsHandler.Delete)
}
