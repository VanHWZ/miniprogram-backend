package representation

import (
	"github.com/gin-gonic/gin"
	repo "miniprogram-backend/repository"
	"time"
)

type Message struct {
	ID        uint           `json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"update_at"`
	Content   string         `json:"content"`

	GroupID    uint          `json:"group"`
	Creator    interface{}   `json:"creator"`
}

type messageSerializer struct {}

var (
	MessageSerializer messageSerializer
)

func init() {
	MessageSerializer = messageSerializer{}
}

func (serializer *messageSerializer) Serialize(message *repo.Message) *Message {
	return &Message{
		ID:        message.ID,
		CreatedAt:  message.CreatedAt,
		UpdatedAt:  message.UpdatedAt,
		Content:   message.Content,
		GroupID:   message.GroupID,
		Creator:   gin.H{"id": message.CreatorID, "name": message.Creator.Name},
	}
}

func (serializer *messageSerializer) SerializeMessages(messages *[]repo.Message) *[]Message {
	var messagesSeriazlized = make([]Message, 0)
	for _, m := range *messages {
		messagesSeriazlized = append(messagesSeriazlized, *serializer.Serialize(&m))
	}
	return &messagesSeriazlized
}

