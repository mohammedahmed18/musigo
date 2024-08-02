package router

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/mohammedahmed18/music-player-rooms/frontend"
	"github.com/mohammedahmed18/music-player-rooms/internal/ws"
	"github.com/rs/zerolog/hlog"
	"github.com/rs/zerolog/log"
)

func Router() *mux.Router {

	router := mux.NewRouter()
	router.Use(hlog.AccessHandler(accessLogger))

	router.HandleFunc("/ws", ws.HandleWebSocket)
	// router.HandleFunc("/room/{roomId}", handleRoomInfo)
	frontend.Register(router)
	return router
}

func accessLogger(r *http.Request, status, size int, dur time.Duration) {
	log.Debug().
		Str("host", r.Host).
		Int("status", status).
		Int("size", size).
		Str("ip", r.RemoteAddr).
		Str("path", r.URL.Path).
		Str("duration", dur.String()).
		Msg("HTTP")
}
