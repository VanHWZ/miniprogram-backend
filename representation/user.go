package representation

import (
	repo "miniprogram-backend/repository"
	"time"
)

type User struct {
	ID uint            `json:"id"`
	Name   string      `json:"name"`
	OpenID string      `json:"open_id"`
	Token string       `json:"token"`
	Group []uint       `json:"group"`
	UpdateAt time.Time `json:"update_at"`
}

type userSerializer struct {}

var (
	UserSerializer userSerializer
)

func init()  {
	UserSerializer = userSerializer{}
}

func (serializer *userSerializer) Serialize(user *repo.User) *User {
	var groupIDs []uint
	for _, g := range user.Groups {
		groupIDs = append(groupIDs, g.ID)
	}
	return &User{
		ID:       user.ID,
		Name:     user.Name,
		OpenID:   user.OpenID,
		Token:    user.Token,
		Group:    groupIDs,
		UpdateAt: user.UpdatedAt,
	}
}

func (serializer *userSerializer) SerializeUsers(users *[]repo.User) *[]User {
	var usersSerialized = make([]User, 0)
	for _, u := range *users {
		usersSerialized = append(usersSerialized, *UserSerializer.Serialize(&u))
	}
	return &usersSerialized
}