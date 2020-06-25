package main

import (
	"fmt"
	"github.com/jakha/chat-go/pkg/websocket"
	"net/http"
)

func main() {
	setupRoutes()
	http.ListenAndServe(":8080", nil)
}

func setupRoutes() {
	pool := websocket.NewPool()
	go pool.Start()

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(pool, w, r)
	})
}

func serveWs(pool *websocket.Pool, w http.ResponseWriter, r *http.Request) {
	fmt.Println("Websocket Endpoint Hit")
	conn, err := websocket.Upgrade(w, r)

	if err != nil {
		fmt.Fprintf(w, "%+v\n", err)
	}

	client := &websocket.Client{
		Conn: conn,
		Pool: pool,
	}

	pool.Register <- client
	client.Read()
}
