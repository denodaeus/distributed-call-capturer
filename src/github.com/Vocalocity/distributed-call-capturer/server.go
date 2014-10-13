package main

import "fmt"

// Server is the server daemon used to query current running jobs and submit new capture jobs.
// The Server is also the entity that capture clients register with to keep distributed tasks in sync.
type Server struct{}

//init creates a new server agent to listen on default configuration settings
func (s Server) init() {
	fmt.Println("Initializing server daemon for RPC communication ...")
}
