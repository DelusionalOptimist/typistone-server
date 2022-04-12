package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/DelusionalOptimist/typistone-server/router"
	"github.com/gorilla/websocket"
)

var (
	port = 8080
	host = "localhost"
)

var upgrader = &websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func main() {
	logger := log.New(os.Stdout, "Typistone-server: ", 0)

	logger.Println("Registering routes...")
	router := router.NewRouter(upgrader, logger)

	address := fmt.Sprint(host, ":", port)

	logger.Printf("Starting listening on %s ...\n", address)
	err := http.ListenAndServe(address, router)
	if err != nil {
		logger.Fatalln(err)
	}
}
