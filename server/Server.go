package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

type Message struct {
	From    string `json:"from"`
	To      string `json:"to"`
	Content string `json:"content"`
}

type Client struct {
	conn     chan Message
	username string
}

var (
	register   = make(chan *Client)
	unregister = make(chan *Client)
	messages   = make(chan Message)
	clients    = make(map[*Client]bool)
)

func handleConnections(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Query().Get("username")

	conn := make(chan Message)
	client := &Client{conn: conn, username: username}

	register <- client
	defer func() { unregister <- client }()

	for {
		var message Message
		err := json.NewDecoder(r.Body).Decode(&message)
		if err != nil {
			log.Println("error decoding message:", err)
			continue
		}

		message.From = username
		messages <- message
		break
	}
}

func handleMessages() {
	for {
		select {
		case client := <-register:
			clients[client] = true
		case client := <-unregister:
			if _, ok := clients[client]; ok {
				close(client.conn)
				delete(clients, client)
			}
		case message := <-messages:
			for client := range clients {
				if client.username == message.To {
					client.conn <- message
				}
			}
		}
	}
}

func main() {
	http.HandleFunc("/ws", handleConnections)

	go handleMessages()

	server := &http.Server{Addr: ":8887"}

	go func() {
		fmt.Println("Server is starting...")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Error starting server: %v", err)
		}
	}()

	waitForShutdown(server)
}

func waitForShutdown(server *http.Server) {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop

	fmt.Println("\nServer is shutting down...")

	if err := server.Shutdown(nil); err != nil {
		log.Fatalf("Error stopping server: %v", err)
	}
}
