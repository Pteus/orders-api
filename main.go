package main

import (
	"context"
	"fmt"
	"github-com/pteus/orders-api/application"
	"os"
	"os/signal"
)

func main() {
	app := application.NewApp()

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop() // defers the execution of stop() until the surrounding function returns

	err := app.Start(ctx)
	if err != nil {
		fmt.Println("failed to start app:", err)
	}
}
