package main

import (
	"log"

	"github.com/horri1520/hori-api/config"
)

func main() {
	e := config.NewEnvVariables()
	if err := e.Init(); err != nil {
		log.Fatal(err)
	}

	s := NewServer()
	if err := s.Init(e); err != nil {
		log.Fatal(err)
	}

	s.Run(e.Port)
}
