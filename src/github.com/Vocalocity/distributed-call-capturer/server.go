package main

import "fmt"

type Server struct{}

func (s Server) init() {
	fmt.Println("Initializing server daemon for RPC communication ...")
}
