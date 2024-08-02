package server

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Server struct {
	router *mux.Router
}

func New(router *mux.Router) (*Server, error) {

	server := &Server{
		router: router,
	}
	return server, nil
}

func (srv *Server) Start(addr string) error {
	return http.ListenAndServe(addr, srv.router)
}

// func handleRoomInfo(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	roomId := vars["roomId"]
// 	info, err := room.GetRoomInfo(roomId)
// 	if err != nil {
// 		b, _ := json.Marshal(err.Error())
// 		w.Write(b)
// 		return
// 	}
// 	b, _ := json.Marshal(info)
// 	w.Write(b)

// }
