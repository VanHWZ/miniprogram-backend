package representation

import (
	repo "miniprogram-backend/repository"
)

type Group struct {
	ID     uint         `json:"id"`
	Name   string       `json:"name"`
	Users  []uint       `json:"users"`
	Events []repo.Event `json:"events"`
}

type groupSerializer struct {}

var (
	GroupSerializer groupSerializer
)

func init()  {
	GroupSerializer = groupSerializer{}
}

func (serializer *groupSerializer) Serialize(groupModel *repo.Group) *Group {
	var userIDs []uint
	for _, u := range groupModel.Users {
		userIDs = append(userIDs, u.ID)
	}
	return &Group{
		ID: groupModel.ID,
		Name: groupModel.Name,
		Users: userIDs,
		Events: groupModel.Events,
	}
}

