package main

import "fmt"
import "flag"

import "github.com/akrennmair/gopcap"

func main() {
	var mode = flag.String("mode", "client", "Mode of execution:  server (handle commands) or client (capture data)")
	flag.Parse()
	fmt.Println("mode", *mode)
}
