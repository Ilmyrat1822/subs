package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/Ilmyrat1822/subs/cmd"
	_ "github.com/Ilmyrat1822/subs/docs"
	internal "github.com/Ilmyrat1822/subs/internal/modules"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title Subscriptions API
// @version 1.0
// @description Subscription management service
// @host localhost:7777
// @BasePath /
func main() {
	server := cmd.NewServer()
	server.Echo.GET("/swagger/*any", echoSwagger.EchoWrapHandler())
	internal.InitRouters(server)

	signalCtx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()
	// Start server
	go func() {
		if err := server.Echo.Start(fmt.Sprintf("0.0.0.0:%v", server.Config.Port)); err != nil && !errors.Is(
			err,
			http.ErrServerClosed,
		) {
			server.Echo.Logger.Fatal("shutting down the server. error: " + err.Error())
		}
	}()

	// Wait for interrupt signal to gracefully shut down the server with a timeout of 10 seconds.
	<-signalCtx.Done()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := server.Echo.Shutdown(ctx); err != nil {
		server.Echo.Logger.Fatal(err)
	}
}
