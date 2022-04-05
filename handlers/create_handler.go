package handlers

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/DelusionalOptimist/typistone-server/models"
	"github.com/gorilla/websocket"
)

// Creates a new lobby.
// Lobby is a waiting area.
// A lobby enters a match when number of players is equal to the minimum
// number set.
func (h *Handler) CreateNewLobby(w http.ResponseWriter, r *http.Request) {
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

	// read lobby config sent by host player
	conn.ReadJSON(lobby)

	lobby.NumberOfPlayers++
	lobby.HostPlayerID = fmt.Sprintf("Player-%d", lobby.NumberOfPlayers)
	lobby.PlayerIDs = append(lobby.PlayerIDs, lobby.HostPlayerID)

	conn.WriteJSON(lobby)
	h.Lobbies[lobby.LobbyID] = lobby

	for h.Lobbies[lobby.LobbyID].NumberOfPlayers < lobby.LobbySize {
		time.Sleep(time.Second * 5)
		conn.WriteMessage(websocket.TextMessage, []byte("Waiting for players"))
	}

	conn.WriteMessage(websocket.TextMessage, []byte("Coolio"))

	conn.WriteJSON(lobby)


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
