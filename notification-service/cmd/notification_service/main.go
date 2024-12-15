package main

import (
	"context"
	"notification-service/internal/app"
	"os"

	"github.com/rs/zerolog/log"
)

func main() {
	ctx := context.Background()

	a, err := app.NewApp(ctx)
	if err != nil {
		log.Error().Err(err).Msg("failed to init app")
		os.Exit(1)
	}

	err = a.Run(ctx)
	if err != nil {
		log.Error().Err(err).Msg("failed to run app")
		os.Exit(1)
	}
}
