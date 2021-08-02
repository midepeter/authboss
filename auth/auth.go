//This is to setup or primary implmentation of th authboss package
package auth

import (
	"log"


	mux "github.com/gorilla/mux"
	"gorm.io/gorm"
	abclientstate "github.com/volatiletech/authboss-clientstate"
	authboss "github.com/volatiletech/authboss/v3"
	_ "github.com/volatiletech/authboss/v3/auth"
	_ "github.com/volatiletech/authboss/v3/logout"
	_ "github.com/volatiletech/authboss/v3/register"
	"github.com/volatiletech/authboss/v3/defaults"
)

type Server struct {
	router *mux.NewRouter
	DB     *gorm.DB
	auth   *authboss.Authboss
}


var (
	ab           = authboss.New()
	database     = gorm.DB
)

func SetUpAuthboss() {
	log.Println("Setting up authentication.....")

	ab.Config.Paths.Mount = "/auth"
	ab.Config.Paths.RootURL = "http://localhost:3000"

	ab.Config.Storage.Server = database

	defaults.SetCore(&ab.Config, false, false)

	if err := ab.init(); err != nil {
		log.Fatalf("Error while initialising authboss -> %s", err)
	}
	//Mounting the router to a path( this should be the same as the path above)
}
