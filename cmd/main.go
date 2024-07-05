package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	track "timeTracker"
	"timeTracker/pkg/handler"
	"timeTracker/pkg/repository"
	"timeTracker/pkg/service"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

// @title TimeTracker App API
// @version 1.0
// @description API Server for TimeTracker Application

// @host localhost:8080
// @BasePath /

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading env variables %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Username: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("SSLMODE"),
	})
	if err != nil {
		logrus.Fatalf("db error connect: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(track.Server)

	if err := srv.Run(os.Getenv("PORT"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
	go func() {
		if err := srv.Run(os.Getenv("PORT"), handlers.InitRoutes()); err != nil {
			log.Fatalf("error occured while running http server: %s", err.Error())
		}
	}()

	logrus.Print("TodoApp Started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Print("TodoApp Shutting Down")

	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error occured on server shutting down: %s", err.Error())
	}

	if err := db.Close(); err != nil {
		logrus.Errorf("error occured on db connection close: %s", err.Error())
	}
}
