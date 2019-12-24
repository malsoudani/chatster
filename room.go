package main

type room struct {
	// forward is a channel that holds incoming messages that will go to other clients
	forwardMsg chan []byte
	join       chan *client
	leave      chan *client
	clients    map[*client]bool
}
