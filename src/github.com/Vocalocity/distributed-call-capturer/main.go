package main

import "fmt"
import "flag"

func main() {
	var mode = flag.String("mode", "client", "Mode of execution:  server (handle commands) or client (capture data)")
	flag.Parse()
	fmt.Println("mode", *mode)

	switch *mode {
	case "client":
		fmt.Println("initializing " + *mode)
	case "server":
		fmt.Println("initializing " + *mode)
		s := Server{}
		s.init()
	default:
		panic("Unable to initialize for mode " + *mode)
	}
}
