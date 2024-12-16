package main

import (
	"context"
	"github.com/rs/zerolog/log"
	"user-service/internal/app"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {
	ctx := context.Background()

	a, err := app.NewApp(ctx)
	if err != nil {
		log.Error().Err(err).Msg("failed to init app")
		return
	}

	err = a.Run()
	if err != nil {
		log.Error().Err(err).Msg("failed to run app")
		return
	}
}
