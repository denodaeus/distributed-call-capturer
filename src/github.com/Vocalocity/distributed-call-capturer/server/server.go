package server

import (
	"fmt"
	"github.com/gorilla/rpc"
	"github.com/gorilla/rpc/json"
	"log"
	"net/http"
)

// Server is the server daemon used to query current running jobs and submit new capture jobs.
// The Server is also the entity that capture clients register with to keep distributed tasks in sync.
type Server struct{}

//Init creates a new server agent to listen on default configuration settings
func (s Server) Init() {
	log.Println("Initializing server daemon for RPC communication ...")

	r := rpc.NewServer()
	r.RegisterCodec(json.NewCodec(), "application/json")
	r.RegisterService(new(APIService), "")
	http.Handle("/rpc", r)
	http.ListenAndServe("localhost:8080", nil)
}

// APIReply is the API reply object
type APIReply struct {
	Message string
}

// RegisterResponse is the server response to register request
type RegisterResponse struct {
	Message string
}

// APIService is the service definition
type APIService struct{}

// RegisteredClient is the exported client registration
type RegisteredClient struct {
	Hostname string
	Role     string
}

var clients = make([]RegisteredClient, 0)

// RegisterClient registers this client to the RPC server
func (a *APIService) RegisterClient(r *http.Request, args *RegisteredClient, reply *APIReply) error {
	hostname := args.Hostname
	role := args.Role
	clients = append(clients, RegisteredClient{Hostname: hostname, Role: role})
	message := "Registered client: " + hostname + " with role " + role
	reply.Message = message
	log.Println(message)
	return nil
}

// ListArgs args count for ListClients
type ListArgs struct {
	Count string
}

// ListClients lists registered clients
func (a *APIService) ListClients(r *http.Request, args *ListArgs, reply *APIReply) error {
	message := fmt.Sprintf("%v", clients)
	log.Println(message)
	reply.Message = message
	return nil
}
