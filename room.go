package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

// in the room struct we hold the messages that we want to send to the clients in a channel to forward the msg
// we have the current state of our current clients of the room in a clients map
// we manage joining and leaving that room through the leave and join channels for their respective purposes
// one of the features that come with using different channels for joining and leaving the client map is that each of them is a different pipeline which means that we won't be accessing the same data resulting in safe opreations without having to worry about corrupt source of truth
type room struct {
	// forward is a channel that holds incoming messages that will go to other clients
	forwardMsg chan []byte
	// channel for clients joining the room
	join chan *client
	// channel for clients leaving the room
	leave chan *client
	// map of the current clients of this room
	clients map[*client]bool
}

const (
	socketBufferSize  = 1024
	messageBufferSize = 256
)

var upgrader = &websocket.Upgrader{ReadBufferSize: socketBufferSize, WriteBufferSize: messageBufferSize}

func (r *room) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	socket, err := upgrader.Upgrade(w, req, nil)
	if err != nil {
		log.Fatal("room server error: ", err)
		return
	}
	client := &client{
		socket:  socket,
		sendMsg: make(chan []byte, messageBufferSize),
		room:    r,
	}
	r.join <- client
	defer func() { r.leave <- client }()
	go client.write()
	client.read()
}

func (r *room) run() {
	for {
		select {
		case client := <-r.join:
			r.clients[client] = true
		case client := <-r.leave:
			// to leave we delete the client from the client list and we close their message sending channel
			delete(r.clients, client)
			close(client.sendMsg)
		case msg := <-r.forwardMsg:
			for client := range r.clients {
				client.sendMsg <- msg
			}
		}
	}
}
