// Package main is main program for the backend-test-golang service.
package main

import (
	"context"
	"os"
	"os/signal"

	"github.com/KryptoKnight/backend-test-golang/api/server"
	"github.com/KryptoKnight/backend-test-golang/config"
	"github.com/KryptoKnight/backend-test-golang/mongo"
	"github.com/rs/zerolog"
)

const (
	configFilenameEnv = "BACKEND_TEST_GOLANG_CFG_ENV"
)

func main() {
	ctx := context.Background()

	osSig := make(chan os.Signal, 1)
	signal.Notify(osSig, os.Interrupt)

	logger := zerolog.New(os.Stdout)

	cfgFileName := os.Getenv(configFilenameEnv)
	if cfgFileName == "" {
		logger.Fatal().Msgf("please define configuration file in %s environment variable", configFilenameEnv)
	}

	appCfg, err := config.NewReader[config.App](cfgFileName).Read()
	if err != nil {
		logger.Fatal().Err(err).Msg("unable to load configurations")
	}

	logger.Info().Msgf("Configurations loaded: %+v", appCfg)

	err = mongo.Connect(ctx, appCfg.DB)
	if err != nil {
		logger.Fatal().Err(err).Msg("unable to connect to the database")
	}

	ser := server.NewServer()
	ser.Configure(appCfg, logger)
	ser.Start()

	<-osSig

	mongo.Close(ctx)
}
