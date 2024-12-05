package main

import (
	"context"
	"fmt"
	"user-service/internal/app"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {
	ctx := context.Background()

	a, err := app.NewApp(ctx)
	if err != nil {
		panic(fmt.Errorf("failed to init app: %s", err.Error()))
	}
	// и здесь про паники то же самое 
	err = a.Run()
	if err != nil {
		panic(fmt.Errorf("failed to run app: %s", err.Error()))
	}
}
