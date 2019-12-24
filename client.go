package main

import "github.com/gorilla/websocket"

type client struct {
	// socket is the websocket for the client
	socket *websocket.Conn
	// send is the channel on which the messages are sent
	sendMsg chan []btye
	// room is the room this client is chatting in
	room *room
}

func (c *client) read() {
	defer c.socket.Close()
	for {
		_, msg, err := c.socket.ReadMessage()
		if err != nil {
			return
		}
		c.room.forwardMsg <- msg
	}
}

func (c *client) write() {
	defer c.socket.Close()
	for msg := range c.sendMsg {
		err := c.socket.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			return
		}
	}
}
