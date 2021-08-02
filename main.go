package main

import (
	"log"
	"net/http"
	"os"

	"github.com/midepeter/authboss/auth"
	"github.com/midepeter/authboss/model"
	"github.com/midepeter/authboss/storer"
)

//This is simple library to implement the use of shim

func (s *Server) LoadRoutes() {
	s.router.Use(s.auth.LoadClientMiddleware)
	s.router.Use(s.redirectIfLoggedIn)

	s.router.Mount("/auth", http.StripPrefix("/auth", s.auth.config.Core.Router))

	log.Println(http.ListenAndServe(":80", s.router))

}

func main() {

	//Initializing the authboss setup function
	ab := auth.SetUpAuthboss()

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
}
