package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/DelusionalOptimist/typistone-server/models"
	clientModels "github.com/DelusionalOptimist/typistone/models"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

// Adds players to the lobbies
func (h *Handler) JoinLobby(w http.ResponseWriter, r *http.Request) {
	// get lobby id from the url
	vars := mux.Vars(r)
	lobbyID, err := strconv.Atoi(vars["lobby_id"])
	if err != nil {
		log.Println(err)
		return
	}

	// find the requested lobby/ TODO: make storage of global state better
	lobby, ok := h.Lobbies[lobbyID]
	if !ok {
		errStr := fmt.Sprintf("Lobby %d doesn't exist", lobbyID)
		http.Error(w, errStr, http.StatusBadRequest)
		return
	}

	if lobby.NumberOfPlayers == lobby.LobbySize {
		errStr := fmt.Sprintf("Unable to join lobby %d because its full.", lobbyID)
		json.NewEncoder(w).Encode(map[string]string{"error": errStr})
		return
	}

	// create a connection with this player
	conn, err := h.Upgrader.Upgrade(w, r, nil)
	if err != nil {
		h.Log.Println(err)
		return
	}

	// initialize player
	player := &models.Player{}
	lobby.NumberOfPlayers++
	player.PlayerID = lobby.NumberOfPlayers
	player.Connection = conn

	// update lobby
	lobby.Players = append(h.Lobbies[lobbyID].Players, player)

	// update global state of the lobby
	h.Lobbies[lobbyID] = lobby

	// send the player info about themselves
	conn.WriteJSON(player)

	for h.Lobbies[lobby.LobbyID].NumberOfPlayers < lobby.LobbySize {
		time.Sleep(time.Second * 5)
		err := conn.WriteMessage(websocket.TextMessage, []byte("Waiting for players"))
		if err != nil {
			log.Println(err)
			return
		}
	}

	conn.WriteMessage(websocket.TextMessage, []byte("Starting game"))
	conn.WriteMessage(websocket.TextMessage, []byte("The quick brown fox jumps over the lazy black dogs."))

	gameData := &clientModels.Game{}
	for {
		err = conn.ReadJSON(gameData)
		if err != nil {
			fmt.Println(err)
		}
		lobby.Game.Percentages[player.PlayerID-1] = 10
	}
}
