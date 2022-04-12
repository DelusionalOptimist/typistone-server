package handlers

import (
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/DelusionalOptimist/typistone-server/models"
	clientModels "github.com/DelusionalOptimist/typistone/models"
	"github.com/gorilla/websocket"
)

// Creates a new lobby.
// Lobby is a waiting area.
// A lobby enters a match when number of players is equal to the minimum
// number set.
func (h *Handler) CreateNewLobby(w http.ResponseWriter, r *http.Request) {
	// create a new lobby and alot it a random number
	lobby := &models.Lobby{}
	rand.Seed(time.Now().Unix())
	lobbyID := rand.Intn(1000)
	lobby.LobbyID = lobbyID

	// connect with the host player
	conn, err := h.Upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	// creates a player object which stores the connection
	player := &models.Player{}
	player.Connection = conn

	// read lobby config sent by host player
	lobbyConfig := &models.Lobby{}
	err = conn.ReadJSON(lobbyConfig)
	if err != nil {
		log.Printf("Error while reading lobby config: %s", err.Error())
		return
	}

	// assignes a number to this player. It would be always 1 for host
	lobby.LobbySize = lobbyConfig.LobbySize
	lobby.NumberOfPlayers++
	lobby.HostPlayerID = lobby.NumberOfPlayers
	player.PlayerID = lobby.NumberOfPlayers
	lobby.Players = append(lobby.Players, player)

	log.Printf("Lobby created. \nID: %d Host ID: %d Size: %d", lobby.LobbyID, player.PlayerID, lobby.LobbySize)

	// send lobby config to the host player
	conn.WriteJSON(lobby)

	// send player info to the host player
	conn.WriteJSON(player)

	// save the created lobby to global lobby collection
	h.Lobbies[lobby.LobbyID] = lobby

	// keep waiting for players till number of paritcipating players is lesser
	// than the lobby size
	for h.Lobbies[lobby.LobbyID].NumberOfPlayers < lobby.LobbySize {
		time.Sleep(time.Second * 5)
		err := conn.WriteMessage(websocket.TextMessage, []byte("Waiting for players"))
		if err != nil {
			log.Println(err)
			return
		}
	}

	// send text and start game
	conn.WriteMessage(websocket.TextMessage, []byte("Starting game"))
	conn.WriteMessage(websocket.TextMessage, []byte("The quick brown fox jumps over the lazy black dogs."))

	// start exchanging data b/w the clients once lobby full

	//gameDataReceived := make(chan clientModels.Game)
	//go func() {
	//	for {
	//		time.Sleep(1 * time.Second)
	//		conn.ReadJSON(gameData)
	//		gameDataReceived <- *gameData
	//	}
	//}()

//	gameDataSent := make(chan clientModels.Game)
//	go func() {
//	}()

	// send state of other players to the host player in every 2 seconds
	gameData := &clientModels.Game{}
	lobby.Game = gameData
	for {
		// get percentages for all players
		for i, p := range h.Lobbies[lobby.LobbyID].Game.Percentages {
			if i == player.PlayerID-1 {
				continue
			}
			gameData.Percentages[i] = p
		}

		time.Sleep(2 * time.Second)
		err := conn.WriteJSON(gameData)
		if err != nil {
			log.Println(err)
			return
		}
	}

	// debug
	//conn.WriteJSON(lobby)

	// get game config from the request body
	//defer r.Body.Close()
	//json.NewDecoder(r.Body).Decode(lobby)

	////timer := time.NewTimer(time.Duration(gameCfg.Timeout) * time.Second)
	//rand.Seed(time.Now().UnixNano())

	//json.NewEncoder(w).Encode(lobby)

	// client now sends an http get to the particular thing

	//for {
	//	player := &models.Player{}
	//	player.LobbyID = lobbyID
	//	conn, err := h.Upgrader.Upgrade(w, r, nil)
	//	if err != nil {
	//		log.Println(err)
	//		return
	//	}
	//	conn.WriteMessage(websocket.TextMessage, []byte("Hell"))
	//}

	// timeout stuff
	//	select {
	//		case <- timer.C:
	//			h.Log.Println("Timed out waiting for other players.")
	//			http.Error(w, "Timed out waiting for other players", http.StatusRequestTimeout)
	//			return
	//		default:
	//	}

}
