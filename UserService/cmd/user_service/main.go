package main

import (
	"context"
	"fmt"
	"user-service/internal/app"
)

func main() {
	ctx := context.Background()

	a, err := app.NewApp(ctx)
	if err != nil {
		panic(fmt.Errorf("failed to init app: %s", err.Error()))
	}

	err = a.Run()
	if err != nil {
		panic(fmt.Errorf("failed to run app: %s", err.Error()))
	}
}
