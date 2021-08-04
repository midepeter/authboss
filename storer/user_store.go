package storer

import (
	"context"
	"log"

	userd "github.com/midepeter/authboss/user"
	authboss "github.com/volatiletech/authboss/v3"
	"gorm.io/gorm"
)

//This is to implement the user interface
type UserStore struct {
	db *gorm.DB
}

func NewUserStore(db *gorm.DB) *UserStore {
	return &UserStore{
		db: db,
	}
}

func (u *UserStore) Load(_ context.Context, key string) (*authboss.User, error) {
	return nil, nil
}

func (u *UserStore) Save(_ context.Context, user authboss.User) error {
	m := user.(*userd.UserValues)

	u.db.Save(&m)

	return nil
}

func NewUser() (u authboss.User) {
	return userd.NewUserValues()
}

func (u *UserStore) Create(user authboss.User) error {
	m := user.(*userd.UserValues)

	if err := u.db.Create(&m).Error; err != nil {
		return err
	}
	log.Println("The user was created")
	return nil
}
