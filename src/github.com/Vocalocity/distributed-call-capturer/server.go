package main

import "fmt"

/* Server daemon type */
type Server struct{}

func (s Server) init() {
	fmt.Println("Initializing server daemon for RPC communication ...")
}
