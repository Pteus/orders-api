package application

import (
	"context"
	"fmt"
	"github-com/pteus/orders-api/middleware"
	"net/http"
)

type App struct {
	router http.Handler
}

func NewApp() *App {
	return &App{
		router: loadRoutes(),
	}
}

func (a *App) Start(ctx context.Context) error {
	server := &http.Server{
		Addr:    ":3000",
		Handler: middleware.Logging(a.router),
	}

	err := server.ListenAndServe()
	if err != nil {
		return fmt.Errorf("could not start server: %w", err)
	}

	return nil
}
