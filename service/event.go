package service

import (
	"fmt"
	"gorm.io/gorm/clause"
	"miniprogram-backend/errorcode"
	repo "miniprogram-backend/repository"
	repr "miniprogram-backend/representation"
	"miniprogram-backend/util"
	"time"
)

type EventRl struct {
	GroupUserRl
	EventID      uint   `uri:"eid" binding:"required"`
}

type EventCtn struct {
	Content   string         `json:"content"`
	EventType int            `json:"event_type" binding:"oneof=0 1"`
	EventTime util.Datetime  `json:"event_time"`
}

func ListEvent(rl GroupUserRl) (*[]repr.Event, *repr.AppError) {
	var events []repo.Event
	if r := repo.DB.Model(&repo.Event{}).Preload("Creator").Preload("Updater").
		Where("group_id=?", rl.GroupID).Find(&events); r.Error != nil {
			return nil, &repr.AppError{Code: errorcode.DatabaseError, Message: r.Error.Error()}
	}
	return repr.EventSerializer.SerializeEvents(&events), nil
}

func CreateEvent(ctn *EventCtn, rl *GroupUserRl) (*repr.Event, *repr.AppError) {
	event := &repo.Event{
		Content:   ctn.Content,
		EventType: ctn.EventType,
		EventTime: time.Time(ctn.EventTime),
		GroupID:   rl.GroupID,
		CreatorID: rl.UserID,
		UpdaterID: rl.UserID,
	}
	if r := repo.DB.Create(event); r.Error != nil {
		return nil, &repr.AppError{Code: errorcode.DatabaseError, Message: r.Error.Error()}
	}
	return repr.EventSerializer.Serialize(event), nil
}

func RetrieveEvent(rl *EventRl) (*repr.Event, *repr.AppError) {
	var event repo.Event
	if r := repo.DB.Preload("Creator").Preload("Updater").
			Where("group_id=?", rl.GroupID).First(&event, rl.EventID); r.Error != nil {
		return nil, &repr.AppError{Code: errorcode.DatabaseError, Message: r.Error.Error()}
	}
	return repr.EventSerializer.Serialize(&event), nil
}

func UpdateEvent(ctn *EventCtn, rl *EventRl) (*repr.Event, *repr.AppError) {
	var event = repo.Event{ID: rl.EventID}
	if err := repo.DB.Model(&event).Association("Updater").Replace(&repo.User{ID: rl.UserID}); err != nil {
		return nil, &repr.AppError{Code: errorcode.DatabaseError, Message: err.Error()}
	}
	if r := repo.DB.Model(&event).Updates(repo.Event{
			Content:   ctn.Content,
			EventType: ctn.EventType,
			EventTime: time.Time(ctn.EventTime)}); r.Error != nil {
		return nil, &repr.AppError{Code: errorcode.DatabaseError, Message: r.Error.Error()}
	}
	return repr.EventSerializer.Serialize(&event), nil
}

func DeleteEvent(rl *EventRl) *repr.AppError {
	var event repo.Event
	if r := repo.DB.Clauses(clause.Returning{}).
			Where(repo.Event{ID: rl.EventID, GroupID: rl.GroupID}).
			Delete(&event); r.Error != nil {
		return &repr.AppError{Code: errorcode.DatabaseError, Message: r.Error.Error()}
	}
	if event.ID == 0 {
		return &repr.AppError{
			Code: errorcode.EventDeleteError,
			Message: fmt.Sprintf("Event(id=%d) not found in Group(id=%d)", rl.EventID, rl.GroupID)}
	}
	return nil
}
