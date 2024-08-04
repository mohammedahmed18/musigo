package main

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"

	"github.com/mohammedahmed18/musigo/internal/config"
	"github.com/mohammedahmed18/musigo/internal/logger"
	"github.com/mohammedahmed18/musigo/internal/router"
	"github.com/mohammedahmed18/musigo/internal/server"
)

func main() {
	// initialize config file
	config.Init("./temp", "cfg") // TODO: use env var for this

	// use error level log in production
	lvl := zerolog.ErrorLevel
	env := viper.GetString("server.env")
	if env == "dev" {
		lvl = zerolog.DebugLevel
	}

	// init logger
	logger.Init(lvl)

	srv, err := server.New(router.Router())
	if err != nil {
		log.Fatal().Msgf("Error creating server: %v", err)
	}

	srvPort := viper.GetString("server.port")

	log.Error().Msg(srv.Start(":" + srvPort).Error())
	os.Exit(1)
}
