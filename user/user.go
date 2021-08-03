package user

import (
	"github.com/midepeter/authboss/model"
)

type UserValues struct {
	Model     *model.User
	arbitrary map[string]string
}

func NewUserValues(user ...*model.User) *UserValues {
	Model := model.User{}
	modelPtr := &Model
	if len(user) > 0 {
		modelPtr = user[0]
	}

	arbitrary := make(map[string]string)

	return &UserValues{
		Model:     modelPtr,
		arbitrary: arbitrary,
	}
}

func (u UserValues) GetPID() (pid string) {
	return u.Model.Email
}

func (u UserValues) PutPID(pid string) {
	u.Model.Email = pid
}
