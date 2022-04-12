package models

import (
	"github.com/gorilla/websocket"
)

type Player struct {
	// this is the player's ID
	PlayerID int `json:"player_id"`

	// websocket connection for this player
	Connection *websocket.Conn `json:"-"`
}
