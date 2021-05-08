package back

import (
	"context"
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
