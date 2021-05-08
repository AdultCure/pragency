package main

import (
	"context"
	"log"
	"net/http"
	"time"
)

// создаем структуру сервера которую будем получать в запуске и остановке сервера

type Server struct {
	httpServer *http.Server
}

// запуск сервера

func (s *Server) Run(port string) error {
	s.httpServer = &http.Server{
		Addr:           ":" + port,
		MaxHeaderBytes: 1 << 20, //one mb
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}

	return s.httpServer.ListenAndServe()
}

// остановка сервера

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}

func main() {
	srv := new(Server)
	if err := srv.Run("8000"); err != nil {
		log.Fatalf("error occured while http server: %s", err.Error())
	}
}
