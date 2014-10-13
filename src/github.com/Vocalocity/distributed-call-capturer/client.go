package main

import "fmt"

type Client struct{}

func (c Client) init() {
	fmt.Println("Initializing client daemon for RPC communication ...")
}
