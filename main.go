package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/midepeter/authboss/auth"
	"github.com/midepeter/authboss/model"
	"github.com/midepeter/authboss/storer"
	authboss "github.com/volatiletech/authboss/v3"
)

//This is simple library to implement the use of shim

func main() {
	//Initializing the authboss setup function
	auth.SetUpAuthboss()

	env := os.Getenv("ENV")
	if env == "" {
		err := godotenv.Load()

		if err != nil {
			log.Fatal("Error loading .env file", err)
		}
	}

	err := storer.SetupPostgres(&storer.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Password: os.Getenv("DB_PASS"),
		User:     os.Getenv("DB_USER"),
	}, &model.User{})

	if err != nil {
		log.Fatal("failed to connect to postgresql database", err)
	}

	if err != nil {
		log.Fatal("failed to setup tables", err)
	}
	log.Println("Auth is running successfully")

	var ab *authboss.Authboss

	router := mux.NewRouter()
	router.Use(ab.LoadClientStateMiddleware)

	router.PathPrefix("/auth").Handler(http.StripPrefix("/auth", ab.Config.Core.Router))

	log.Println(http.ListenAndServe(":3000", router))
}
