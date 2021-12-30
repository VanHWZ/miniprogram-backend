package representation

import (
	"github.com/gin-gonic/gin"
	repo "miniprogram-backend/repository"
	"miniprogram-backend/util"
	"time"
)

type Event struct {
	ID         uint          `json:"id"`
	CreateAt   time.Time     `json:"create_at"`
	UpdateAt   time.Time     `json:"update_at"`
	Content    string        `json:"content"`
	EventType  int           `json:"event_type"`
	EventTime  util.Datetime `json:"event_time"`

	GroupID    uint          `json:"group"`
	Creator    interface{}   `json:"creator"`
	Updater    interface{}   `json:"updater"`
}

type eventSerializer struct {}

var (
	EventSerializer eventSerializer
)

func init() {
	EventSerializer = eventSerializer{}
}

func (serializer *eventSerializer) Serialize(event *repo.Event) *Event {
	return &Event{
		ID: event.ID,
		CreateAt: event.CreatedAt,
		UpdateAt: event.UpdatedAt,
		Content: event.Content,
		EventType: event.EventType,
		EventTime: util.Datetime(event.EventTime),
		GroupID: event.GroupID,
		Creator: gin.H{"id": event.CreatorID, "name": event.Creator.Name},
		Updater: gin.H{"id": event.UpdaterID, "name": event.Updater.Name},
	}
}

func (serializer *eventSerializer) SerializeEvents(events *[]repo.Event) *[]Event {
	var eventsSeriazlized = make([]Event, 0)
	for _, e := range *events {
		eventsSeriazlized = append(eventsSeriazlized, *serializer.Serialize(&e))
	}
	return &eventsSeriazlized
}