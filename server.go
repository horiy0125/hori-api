package main

import "net/http"

type Server struct {
	handler http.Handler
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Init() {
	mux := http.NewServeMux()

	mux.HandleFunc("/ping", s.PingHandler)

	s.handler = mux
}
