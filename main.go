package main

import (
	"fmt"
	"net/http"

	"github.com/horri1520/hori-api/config"
)

func (s *Server) PingHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "pong")
}

func main() {
	e := config.NewEnvVariables()
	e.Init()

	s := NewServer()
	s.Init()

	http.ListenAndServe(":"+e.Port, s.handler)
}
