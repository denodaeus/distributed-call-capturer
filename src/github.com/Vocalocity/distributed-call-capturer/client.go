package main

import (
	"bytes"
	"github.com/gorilla/rpc"
	"github.com/gorilla/rpc/json"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
)

// Client is the daemon representing a capture client.  It receives commands to queue tasks from the running
// server daemon and is configured to register a keep alive mechanism and report to it.
type Client struct {
	Hostname string
	Role     string
}

const DefaultRemoteAddr = "127.0.0.1"

var hostname, _ = os.Hostname()

const role = "controller"

//init creates a new client agent to listen on configured bind settings
func (c Client) init() {
	log.Println("Initializing client daemon for RPC communication ...")
	c.Register()
}

func execute(s *rpc.Server, method string, req, res interface{}) error {
	if !s.HasMethod(method) {
		//t.Fatal("RPC Expected method to be exported: ", method)
	}

	buf, _ := json.EncodeClientRequest(method, req)
	body := bytes.NewBuffer(buf)
	r, _ := http.NewRequest("POST", "http://localhost:8080/", body)
	r.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	s.ServeHTTP(w, r)

	return json.DecodeClientResponse(w.Body, res)
}

func (c *Client) Register() {
	log.Println("Registering client " + c.Hostname + " to server ")

	s := rpc.NewServer()
	s.RegisterCodec(json.NewCodec(), "application/json")
	s.RegisterService(new(ApiService), "")

	var res RegisterResponse
	execute(s, "ApiService.RegisterClient", &RegisteredClient{hostname, role}, &res)
	return
}
