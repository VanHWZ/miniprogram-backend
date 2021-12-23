package representation

import (
	"miniprogram-backend/repository"
	"time"
)

type User struct {
	ID uint            `json:"id"`
	Name string        `json:"name"`
	OpenID string      `json:"open_id"`
	Token string       `json:"token"`
	Group uint         `json:"group"`
	UpdateAt time.Time `json:"update_at"`
}

type userSerializer struct {}

var (
	UserSerializer userSerializer
)

func init()  {
	UserSerializer = userSerializer{}
}

func (user *userSerializer) Serialize(userModel *repository.User) *User {
	return &User{
		ID: userModel.ID,
		Name: userModel.Name,
		OpenID: userModel.OpenID,
		Token: userModel.Token,
		Group: userModel.Group.ID,
		UpdateAt: userModel.UpdatedAt,
	}
}