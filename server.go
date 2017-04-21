package main

import (
	"bustrack/myredis"
	"bustrack/routes"
	"context"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	poolCreatedStatus := myredis.InitPool()
	if poolCreatedStatus == true {
		e := echo.New()
		// Middleware
		e.Use(middleware.Logger())
		e.Use(middleware.Recover())
		//CORS
		e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins: []string{"*"},
			AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
		}))

		// Routes
		routes.Route(e)

		// Server
		go func() {
			if err := e.Start(":1323"); err != nil {
				e.Logger.Info("shutting down the server")
			}
		}()
		// Wait for interrupt signal to gracefully shutdown the server with
		// a ttimeout of 10 seconds.
		quit := make(chan os.Signal)
		signal.Notify(quit, os.Interrupt)
		<-quit
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		if err := e.Shutdown(ctx); err != nil {
			e.Logger.Fatal(err)
		}
	}
	//fmt.Println("initialization pool status:" + poolCreatedStatus)
}
