package main

import (
	"github.com/gorilla/rpc"
	"github.com/gorilla/rpc/json"
	"log"
	"net/http"
)

// Server is the server daemon used to query current running jobs and submit new capture jobs.
// The Server is also the entity that capture clients register with to keep distributed tasks in sync.
type Server struct{}

//init creates a new server agent to listen on default configuration settings
func (s Server) init() {
	log.Println("Initializing server daemon for RPC communication ...")

	r := rpc.NewServer()
	r.RegisterCodec(json.NewCodec(), "application/json")
	r.RegisterService(new(ApiService), "")
	http.Handle("/rpc", r)
	http.ListenAndServe("localhost:8080", nil)
}

type ApiArgs struct {
	Who string
}

type ApiReply struct {
	Message string
}

type ApiService struct{}

func (a *ApiService) Say(r *http.Request, args *ApiArgs, reply *ApiReply) error {
	reply.Message = "Hello, " + args.Who + "!"
	return nil
}
