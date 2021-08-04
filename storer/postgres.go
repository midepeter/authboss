package storer

import (
	"fmt"
	"log"

	authboss "github.com/volatiletech/authboss/v3"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	Host     string
	Port     string
	Password string
	User     string
}

func SetupPostgres(config *Config, models ...interface{}) error {
	var (
		db  *gorm.DB
		err error
	)

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s",
		config.Host, config.Port, config.User, config.Password)

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	if err := db.AutoMigrate(models...); err != nil {
		log.Println("Print error", err)
	}
	return err
}

type postgresStore struct {
	authboss.CreatingServerStorer
	db *gorm.DB
}

func NewpostgresStore(db *gorm.DB) *postgresStore {
	return &postgresStore{
		db: db,
	}
}
