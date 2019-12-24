package main

type room struct {
	// forward is a channel that holds incoming messages that will go to other clients
	forward chan []byte
}
