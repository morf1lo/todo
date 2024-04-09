package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/morf1lo/todo"
	"github.com/morf1lo/todo/internal/config"
	"github.com/morf1lo/todo/pkg/db"
	"github.com/morf1lo/todo/internal/handler"
	"github.com/morf1lo/todo/internal/repository"
	"github.com/morf1lo/todo/internal/service"
)

func main() {
	if err := config.Init(); err != nil {
		log.Fatal(err)
	}

	client, err := db.Connect(context.TODO(), os.Getenv("MONGO_URL"))
	if err != nil {
		log.Fatal(err)
	}

	repos := repository.New(client.Database("todo"))
	services := service.New(repos)
	handlers := handler.New(services)

	server := new(todo.Server)
	go func() {
		if err := server.Run(os.Getenv("PORT"), handlers.InitRoutes()); err != nil {
			log.Fatalf("error occured while running the server: %s\n", err.Error())
		}
	}()

	log.Println("Server Started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	log.Println("Server Shutting Down")

	if err := server.Shutdown(context.Background()); err != nil {
		log.Fatalf("error occured on server shutting down: %s\n", err.Error())
	}

	if err := client.Disconnect(context.Background()); err != nil {
		log.Fatalf("error occured on db connection close: %s\n", err.Error())
	}
}
