package models

type Player struct {
	// this is the player's ID
	PlayerID string `json:"player_id"`

	// this is player's lobby ID
	LobbyID int `json:"lobby_id"`
}
