package server

import (
	"github.com/koding/kite"
	"log"
)

// Server is the server daemon used to query current running jobs and submit new capture jobs.
// The Server is also the entity that capture clients register with to keep distributed tasks in sync.
type Server struct{}

//Init creates a new server agent to listen on default configuration settings
func (s Server) Init() {
	log.Println("Initializing server daemon for RPC communication ...")

	k := kite.New("server", "1.0.0")
	k.Config.Port = 6000
	k.Config.DisableAuthentication = true

	k.HandleFunc("register", func(r *kite.Request) (interface{}, error) {
		a := r.Args.One().MustString()
		log.Println("register :: sending register request, a" + a)
		return a, nil
	})

	k.HandleFunc("start", func(r *kite.Request) (interface{}, error) {
		a := r.Args.One().MustString()
		log.Println("start :: received start for call-id " + a)
		return a, nil
	})

	k.HandleFunc("stop", func(r *kite.Request) (interface{}, error) {
		a := r.Args.One().MustString()
		log.Println("stop :: received stop for call-id" + a)
		return a, nil
	})

	k.HandleFunc("stream", func(r *kite.Request) (interface{}, error) {
		args := r.Args.MustSliceOfLength(2)
		client := args[0].MustString()
		cid := args[1].MustString()
		log.Println("stream :: received stream for call-id" + cid + ", client=" + client)
		return cid, nil
	})

	k.Run()
}
