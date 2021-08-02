package auth

import (
	"github.com/midepeter/authboss/model"
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

func (u *UserStore) Load(key string) (authboss.User, error) {
	var user *model.User
	activeUser := u.db.Where("id = ?", key).Find(&user)
	return activeUser, nil
}

func (u *UserStore) Save(user authboss.User) error {
	return nil
}

func New(user authboss.User) *authboss.User {
	return user
}
func (u *UserStore) Create(user authboss.User) error {

}
