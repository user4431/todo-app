package main

import (
	"github.com/user4431/todo-app"
	"github.com/user4431/todo-app/pkg/handler"
	"github.com/user4431/todo-app/pkg/repository"
	"github.com/user4431/todo-app/pkg/service"
	"log"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)
	if err := srv.Run("9090", handlers.InitRoutes()); err != nil {
		log.Fatal(err)
	}
}
