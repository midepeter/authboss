package storer

import (
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

func (u *UserStore) Load(key string) (*authboss.User, error) {
	return &User, nil
}

func (u *UserStore) Save(user authboss.User) error {
	return nil
}

func NewUser(user authboss.User) *authboss.User {
	return nil
}
func (u *UserStore) Create(user authboss.User) error {
	return nil
}
