package router

import (
	"log"

	"github.com/DelusionalOptimist/typistone-server/handlers"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

func NewRouter(upgrader *websocket.Upgrader, log *log.Logger) *mux.Router {
	r := mux.NewRouter()

	h := handlers.NewHandler(upgrader, log)

	r.HandleFunc("/create", h.CreateNewLobby)
	r.HandleFunc("/join/{lobby_id}", h.JoinLobby)

	return r
}
