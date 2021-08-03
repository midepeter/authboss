//This is to setup or primary implmentation of th authboss package
package auth

import (
	"log"

	storer "github.com/midepeter/authboss/storer"
	authboss "github.com/volatiletech/authboss/v3"
	_ "github.com/volatiletech/authboss/v3/auth"
	"github.com/volatiletech/authboss/v3/defaults"
	_ "github.com/volatiletech/authboss/v3/logout"
	_ "github.com/volatiletech/authboss/v3/register"
	"gorm.io/gorm"
)

type Server struct {
	DB   *gorm.DB
	auth *authboss.Authboss
}

var (
	ab       = authboss.New()
	database = storer.New()
)

func SetUpAuthboss() {
	log.Println("Setting up authentication.....")

	ab.Config.Paths.Mount = "/auth"
	ab.Config.Paths.RootURL = "http://localhost:3000"

	ab.Config.Storage.Server = database

	defaults.SetCore(&ab.Config, false, false)

	if err := ab.Init(); err != nil {
		log.Fatalf("Error while initialising authboss -> %s", err)
	}
	//Mounting the router to a path( this should be the same as the path above)
}
