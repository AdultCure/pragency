package main

import (
	"log"
	backend "pragency"
	"pragency/pkg/handler"
	"pragency/pkg/repository"
	"pragency/pkg/service"
)


func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(backend.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error")
	}
}
