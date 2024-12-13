package main

import (
	"context"
	"fmt"
	"log"
	"user-service/internal/app"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {
	ctx := context.Background()

	a, err := app.NewApp(ctx)
	if err != nil {
		log.Println(fmt.Errorf("failed to init app: %s", err.Error()))
		return
	}

	err = a.Run()
	if err != nil {
		log.Println(fmt.Errorf("failed to run app: %s", err.Error()))
		return
	}
}
