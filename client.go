package main

import "github.com/gorilla/websocket"

type client struct {
	// socket is the websocket for the client
	socket *websocket.Conn
	// send is the channel on which the messages are sent
	send chan []btye
	// room is the room this client is chatting in
	room *room
}
