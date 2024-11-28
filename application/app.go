package application

import (
	"context"
	"fmt"
	"github-com/pteus/orders-api/middleware"
	"github.com/redis/go-redis/v9"
	"net/http"
	"time"
)

type App struct {
	router http.Handler
	rdb    *redis.Client
}

func NewApp() *App {
	return &App{
		router: loadRoutes(),
		rdb:    redis.NewClient(&redis.Options{}),
	}
}

func (a *App) Start(ctx context.Context) error {
	server := &http.Server{
		Addr:    ":3000",
		Handler: middleware.Logging(a.router),
	}

	err := a.rdb.Ping(ctx).Err()
	if err != nil {
		return fmt.Errorf("could not connect to redis: %w", err)
	}

	fmt.Println("starting server...")

	ch := make(chan error, 1)

	go func() {
		err = server.ListenAndServe()
		if err != nil {
			ch <- fmt.Errorf("could not start server: %w", err)
		}
		close(ch)
	}()

	defer func() {
		if err := a.rdb.Close(); err != nil {
			fmt.Println("could not close redis connection:", err)
		}
	}()

	fmt.Println("listening on port 3000")

	select {
	case err := <-ch:
		return err
	case <-ctx.Done():
		timeout, cancelFunc := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancelFunc()

		return server.Shutdown(timeout)
	}
}
