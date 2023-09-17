package main

import (
	"encoding/json"
	"fmt"
	"log"
	nameenrich "name-enrich"
	"name-enrich/pkg/handler"
	"name-enrich/pkg/repository"
	"name-enrich/pkg/service"
	"net/http"
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

	// db, err := repository.NewPostgresDb(repository.Config{
	// 	Host:     os.Getenv("DB_HOST"),
	// 	Port:     os.Getenv("DB_PORT"),
	// 	Username: os.Getenv("DB_USERNAME"),
	// 	Password: os.Getenv("DB_PASSWORD"),
	// 	DBName:   os.Getenv("DBNAME"),
	// 	SSLMode:  os.Getenv("DB_SSLMODE"),
	// })

	// if err != nil {
	// 	logrus.Fatalf("cant init db %s", err)
	// }

	//repo := repository.NewRepository(db)
	repo := repository.NewRepository()
	service := service.NewService(repo)
	handlers := handler.NewHandler(service)

	//http.HandleFunc("/", personEncode)

	srv := new(nameenrich.Server)

	if err := srv.Run(os.Getenv("PORT"), handlers.InitRoutes()); err != nil {
		logrus.Fatalf("error while running http server: %s", err.Error())
	}
}

func readPerson(w http.ResponseWriter, r *http.Request) {
	var p nameenrich.Person

	err := json.NewDecoder(r.Body).Decode(&p)

	if err != nil {
		logrus.Error(err.Error())
		return
	}

	if p.Name == "" || p.Surname == "" {
		logrus.Error("incorrect data")
		return
	}

	urlAge := fmt.Sprintf("https://api.agify.io/?name=%s", p.Name)

	resp, err := http.Get(urlAge)

	if err != nil {
		logrus.Error("faild to connect api " + err.Error())
		return
	}

	err = json.NewDecoder(resp.Body).Decode(&p)

	if err != nil {
		logrus.Error(err.Error())
		return
	}

	urlGender := fmt.Sprintf("https://api.genderize.io/?name=%s", p.Name)

	resp, err = http.Get(urlGender)

	if err != nil {
		logrus.Error("faild to connect api " + err.Error())
		return
	}

	err = json.NewDecoder(resp.Body).Decode(&p)

	if err != nil {
		logrus.Error(err.Error())
		return
	}

	urlNationality := fmt.Sprintf("https://api.nationalize.io/?name=%s", p.Name)

	resp, err = http.Get(urlNationality)

	if err != nil {
		logrus.Error("faild to connect api " + err.Error())
		return
	}

	var c nameenrich.Country

	err = json.NewDecoder(resp.Body).Decode(&c)

	if err != nil {
		logrus.Error(err.Error())
		return
	}

	p.Country = c.Country[0].CountryId
	log.Println(p)
}
