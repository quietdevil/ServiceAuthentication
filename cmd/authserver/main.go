package main

import (
	"context"
	"log"
	"serviceauth/internal/app"
)

func main() {
	ctx := context.Background()

	app, err := app.NewApp(ctx)
	if err != nil {
		log.Fatal(err)
	}

	if err = app.Run(); err != nil {
		log.Fatal(err)
	}
}
