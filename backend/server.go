package backend

import (
	"context"
	"net/http"
	"time"
)

// Объявление структуры сервера из базового пакета net/http

type Server struct {
	httpServer *http.Server
}

// Функция создания сервера

func (s *Server) Run(port string, handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        handler,
		MaxHeaderBytes: 1 << 20,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}

	return s.httpServer.ListenAndServe()
}

// Функция закрытия сервера

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
