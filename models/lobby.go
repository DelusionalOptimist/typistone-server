package models

type LobbyConfig struct {
	// the number of players needed in the lobby before a match can be
	// started
	LobbySize int `json:"lobby_size"`

	// timeout (in seconds) waiting for new members
	Timeout int `json:"timeout"`
}

type Lobby struct {
	*LobbyConfig `json:"lobby_config"`

	// ID of this lobby
	LobbyID int `json:"lobby_id"`

	// number of players present in this lobby
	NumberOfPlayers int `json:"number_of_players"`

	// IDs of the players in this lobby
	PlayerIDs []string `json:"player_ids"`

	// Host player ID
	HostPlayerID string `json:"host_player_id"`
}
