package main

import (
	"context"
	"os"
	"project-service/internal/app"

	"github.com/rs/zerolog/log"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {
	ctx := context.Background()

	a, err := app.NewApp(ctx)
	if err != nil {
		log.Error().Err(err).Msg("failed to init app")
		os.Exit(1)
	}

	err = a.Run()
	if err != nil {
		log.Error().Err(err).Msg("failed to run app")
		os.Exit(1)
	}
}
