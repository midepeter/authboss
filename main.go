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
)

//This is simple library to implement the use of shim

func (s *Server) LoadRoutes() {

	router := mux.NewRouter()
	router.Use(s.auth.LoadClientMiddleware)
	router.Use(s.redirectIfLoggedIn)

	router.Mount("/auth", http.StripPrefix("/auth", s.auth.config.Core.Router))

	log.Println(http.ListenAndServe(":80", router))

}

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

	db, err := storer.New(&storer.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Password: os.Getenv("DB_PASS"),
		User:     os.Getenv("DB_USER"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
	})

	if err != nil {
		log.Fatal("failed to connect to postgresql database", err)
	}

	err = storer.SetupDatabase(db,
		&model.User{},
	)

	if err != nil {
		log.Fatal("failed to setup tables", err)
	}
	log.Println("Auth is running successfully")
}
