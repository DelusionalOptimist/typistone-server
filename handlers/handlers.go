package handlers

import (
	"log"

	"github.com/DelusionalOptimist/typistone-server/models"
	"github.com/gorilla/websocket"
)

type Handler struct {
	Upgrader *websocket.Upgrader
	Log *log.Logger
	Lobbies map[int]*models.Lobby
}

func NewHandler(upgrader *websocket.Upgrader, log *log.Logger) *Handler {
	lobbies := make(map[int]*models.Lobby)

	return &Handler{
		Upgrader: upgrader,
		Log: log,
		Lobbies: lobbies,
	}
}
