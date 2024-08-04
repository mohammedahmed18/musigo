package router

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog/hlog"
	"github.com/rs/zerolog/log"

	"github.com/mohammedahmed18/musigo/frontend"
	"github.com/mohammedahmed18/musigo/internal/ws"
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
