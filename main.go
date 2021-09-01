package main

import (
	"fmt"
	"net/http"
)

type Server struct {
	handler http.Handler
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) init() {
	mux := http.NewServeMux()

	mux.HandleFunc("/ping", s.PingHandler)

	s.handler = mux
}

func (s *Server) PingHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "pong")
}

func main() {
	s := NewServer()
	s.init()

	http.ListenAndServe(":8080", s.handler)
}
