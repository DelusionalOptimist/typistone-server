package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Adds players to the lobbies
func (h *Handler) JoinLobby(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		return
	}

	lobby, ok := h.Lobbies[id]
	if !ok {
		errStr := fmt.Sprintf("Lobby %d doesn't exist", id)
		http.Error(w, errStr, http.StatusBadRequest)
		return
	}

	playerID := fmt.Sprintf("Player-%d", lobby.NumberOfPlayers+1)
	lobby.NumberOfPlayers++

	lobby.PlayerIDs = append(h.Lobbies[id].PlayerIDs, playerID)

	// update global lobby state
	h.Lobbies[id] = lobby

	conn, err := h.Upgrader.Upgrade(w, r, nil)
	if err != nil {
		h.Log.Println(err)
		return
	}

	conn.WriteJSON(lobby)
}
