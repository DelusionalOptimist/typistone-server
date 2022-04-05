package main

import (
	"log"
	"net/http"
	"os"

	"github.com/DelusionalOptimist/typistone-server/router"
	"github.com/gorilla/websocket"
)

var upgrader = &websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func main() {
	logger := log.New(os.Stdout, "Typistone-server: ", 0)
	router := router.NewRouter(upgrader, logger)

	http.ListenAndServe("localhost:8080", router)
}
