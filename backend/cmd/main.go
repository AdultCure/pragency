package main

import (
	"context"
	"log"
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
}

func (s* Server) Run(port string) error {
	s.httpServer = &http.Server{
		Addr: ":" + port,
		MaxHeaderBytes: 1 << 20,
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}

func main() {
	srv := new(Server)
	if err := srv.Run("8000"); err != nil {
		log.Fatalf("error")
	}
}
