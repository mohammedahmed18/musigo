package main

import (
	"context"

	"github.com/mohammedahmed18/music-player-rooms/internal/logger"
	"github.com/mohammedahmed18/music-player-rooms/internal/room"
	"github.com/mohammedahmed18/music-player-rooms/internal/router"
	"github.com/mohammedahmed18/music-player-rooms/internal/server"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	logger.Init(zerolog.DebugLevel) // TODO: user env var for this
	srv, err := server.New(router.Router())
	if err != nil {
		log.Fatal().Msgf("Error creating server: %v", err)
	}

	go func() {
		err = srv.Start(":8080")
		if err != nil {
			log.Fatal().Msgf("Error starting server: %v", err)
		}
	}()

	r := room.New()
	r.ServerURL = "ws://localhost:8080/ws"

	ctx, closeRoom := context.WithCancel(context.Background())
	defer closeRoom()
	go func() {
		r.Start(ctx)
	}()

	// r.PlayMusic("raad elkurdi")
	// r.Join("user_1")
	// r.Join("user_2")
	// r.Join("user_3")
	// r.Join("user_4")
	// r.Join("user_4")
	// r.PlayMusic("raad elkurdi 222")
	// r.Shutdown()

	// Prevent main from exiting
	select {}
}
