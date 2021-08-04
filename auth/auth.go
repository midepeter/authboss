//This is to setup or primary implmentation of th authboss package
package auth

import (
	"log"

	"github.com/midepeter/authboss/storer"
	abrenderer "github.com/volatiletech/authboss-renderer"
	authboss "github.com/volatiletech/authboss/v3"
	_ "github.com/volatiletech/authboss/v3/auth"
	"github.com/volatiletech/authboss/v3/defaults"
	_ "github.com/volatiletech/authboss/v3/logout"
	_ "github.com/volatiletech/authboss/v3/register"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func SetUpAuthboss() {
	log.Println("Setting up authentication")

	ab := authboss.New()

	ab.Config.Paths.Mount = "/auth"
	ab.Config.Paths.RootURL = "http://localhost:3000"

	ab.Config.Core.ViewRenderer = abrenderer.NewHTML("/auth", "ab_views")

	ab.Config.Storage.Server = storer.NewpostgresStore(db)
	ab.Config.Storage.SessionState = nil

	defaults.SetCore(&ab.Config, false, false)

	//Initializing the authboss package
	if err := ab.Init(); err != nil {
		panic(err)
	}

	log.Println("Authentication setup finished")
}
