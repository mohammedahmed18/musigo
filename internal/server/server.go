package server

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"github.com/mohammedahmed18/music-player-rooms/internal/room"
	"github.com/rs/zerolog/log"
)

type Server struct {
	router   *mux.Router
	upgrader websocket.Upgrader
}

func New() (*Server, error) {
	router := mux.NewRouter()
	upgrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true // Allow connections from any origin
		},
	}

	server := &Server{
		router:   router,
		upgrader: upgrader,
	}
	return server, nil
}

func (srv *Server) handleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := srv.upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, "Could not upgrade to websocket", http.StatusInternalServerError)
		return
	}
	defer conn.Close()

	log.Debug().Msg("Client connected")
	for {
		_, _, err := conn.ReadMessage()
		if err != nil {
			break
		}
	}
	log.Debug().Msg("Client disconnected")
}

func (srv *Server) Start(addr string) error {
	srv.routes()
	return http.ListenAndServe(addr, srv.router)
}

func handleRoomInfo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	roomId := vars["roomId"]
	info, err := room.GetRoomInfo(roomId)
	if err != nil {
		b, _ := json.Marshal(err.Error())
		w.Write(b)
		return
	}
	b, _ := json.Marshal(info)
	w.Write(b)

}
func (srv *Server) routes() {
	srv.router.HandleFunc("/ws", srv.handleWebSocket)
	srv.router.HandleFunc("/room/{roomId}", handleRoomInfo)
}
