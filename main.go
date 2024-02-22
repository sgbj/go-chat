package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrade = websocket.Upgrader{}

var (
	clients    = make(map[*websocket.Conn]bool)
	register   = make(chan *websocket.Conn)
	unregister = make(chan *websocket.Conn)
	broadcast  = make(chan []byte)
)

func run() {
	for {
		select {
		case client := <-register:
			clients[client] = true

		case client := <-unregister:
			delete(clients, client)
			client.Close()

		case message := <-broadcast:
			for client := range clients {
				if err := client.WriteMessage(websocket.TextMessage, message); err != nil {
					delete(clients, client)
					client.Close()
				}
			}
		}
	}
}

func chat(w http.ResponseWriter, r *http.Request) {
	c, err := upgrade.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer func() {
		unregister <- c
		c.Close()
	}()

	register <- c

	for {
		_, message, err := c.ReadMessage()
		if err != nil {
			log.Println(err)
			break
		}

		broadcast <- message
	}
}

func main() {
	mux := http.NewServeMux()

	mux.Handle("/", http.FileServer(http.Dir("./static")))
	mux.HandleFunc("/ws", chat)

	go run()

	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
