package main

import (
	nameenrich "name-enrich"
	"name-enrich/pkg/handler"
	"name-enrich/pkg/repository"
	"name-enrich/pkg/service"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

func main() {

	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading env vars %s", err)
	}

	db, err := repository.NewPostgresDb(repository.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Username: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
	})

	if err != nil {
		logrus.Fatalf("cant init db %s", err)
	}

	repo := repository.NewRepository(db)
	service := service.NewService(repo)
	handlers := handler.NewHandler(service)

	srv := new(nameenrich.Server)

	if err := srv.Run(os.Getenv("PORT"), handlers.InitRoutes()); err != nil {
		logrus.Fatalf("error while running http server: %s", err.Error())
	}
}
