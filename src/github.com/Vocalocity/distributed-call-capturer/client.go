package main

import "fmt"

// Client is the daemon representing a capture client.  It receives commands to queue tasks from the running
// server daemon and is configured to register a keep alive mechanism and report to it.
type Client struct{}

//init creates a new client agent to listen on configured bind settings
func (c Client) init() {
	fmt.Println("Initializing client daemon for RPC communication ...")
}
