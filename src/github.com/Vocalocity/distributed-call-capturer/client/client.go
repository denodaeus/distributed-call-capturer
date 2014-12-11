package client

import (
	"fmt"
	"github.com/koding/kite"
	"log"
	"os"
)

// Client is the daemon representing a capture client.  It receives commands to queue tasks from the running
// server daemon and is configured to register a keep alive mechanism and report to it.
type Client struct {
	Hostname string
	Role     string
}

const defaultRemoteAddr = "127.0.0.1"
const rpcPort = "8080"
const rpcURL = "http://" + defaultRemoteAddr + ":" + rpcPort + "/rpc"
const role = "controller"

var hostname, _ = os.Hostname()

//Init creates a new client agent to listen on configured bind settings
func (c Client) Init() {
	log.Println("Initializing client daemon for RPC communication ...")

	k := kite.New("client", "1.0.0")
	client := k.NewClient("http://localhost:6000/kite")
	client.Dial()

	response, _ := client.Tell("register", "controller")
	fmt.Println(response.MustString())

	response, _ = client.Tell("start", "123457@7.8.9.10")
	fmt.Println(response.MustString())

	response, _ = client.Tell("stop", "123457@7.8.9.10")
	fmt.Println(response.MustString())

	response, _ = client.Tell("stream", "localhost:11211", "123457@7.8.9.10")
	fmt.Println(response.MustString())

}

func (c Client) Register() {
}
