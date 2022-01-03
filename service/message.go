package service

import (
	"fmt"
	"gorm.io/gorm/clause"
	"miniprogram-backend/errorcode"
	repo "miniprogram-backend/repository"
	repr "miniprogram-backend/representation"
)

type MessageRl struct {
	GroupUserRl
	MessageID    uint  `uri:"mid" binding:"required"`
}

type MessageCtn struct {
	Content   string  `json:"content"`
}

func ListMessage(rl GroupUserRl) (*[]repr.Message, *repr.AppError) {
	var messages []repo.Message
	if r := repo.DB.Model(&repo.Message{}).Preload("Creator").
		Where("group_id=? AND creator_id=?", rl.GroupID, rl.UserID).
		Find(&messages); r.Error != nil {
		return nil, &repr.AppError{Code: errorcode.DatabaseError, Message: r.Error.Error()}
	}
	return repr.MessageSerializer.SerializeMessages(&messages), nil
}

func CreateMessage(ctn *MessageCtn, rl *GroupUserRl) (*repr.Message, *repr.AppError) {
	message := &repo.Message{
		Content: ctn.Content,
		GroupID: rl.GroupID,
		CreatorID: rl.UserID,
	}
	if r := repo.DB.Create(message); r.Error != nil {
		return nil, &repr.AppError{Code: errorcode.DatabaseError, Message: r.Error.Error()}
	}
	return repr.MessageSerializer.Serialize(message), nil
}

func RetrieveMessage(rl *MessageRl) (*repr.Message, *repr.AppError) {
	var message repo.Message
	if r := repo.DB.Preload("Creator").Where("group_id=?", rl.GroupID).
			First(&message, rl.MessageID); r.Error != nil {
		return nil, &repr.AppError{Code: errorcode.DatabaseError, Message: r.Error.Error()}
	}
	return repr.MessageSerializer.Serialize(&message), nil
}

func UpdateMessage(ctn *MessageCtn, rl *MessageRl) (*repr.Message, *repr.AppError) {
	var message = repo.Message{ID: rl.MessageID}
	if r := repo.DB.Model(&message).Updates(repo.Message{Content: ctn.Content}); r.Error != nil {
		return nil, &repr.AppError{Code: errorcode.DatabaseError, Message: r.Error.Error()}
	}
	return repr.MessageSerializer.Serialize(&message), nil
}

func DeleteMessage(rl *MessageRl) *repr.AppError {
	var message repo.Message
	if r := repo.DB.Clauses(clause.Returning{}).
		Where(repo.Message{ID: rl.MessageID, GroupID: rl.GroupID}).
		Delete(&message); r.Error != nil {
		return &repr.AppError{Code: errorcode.DatabaseError, Message: r.Error.Error()}
	}
	if message.ID == 0 {
		return &repr.AppError{
			Code: errorcode.EventDeleteError,
			Message: fmt.Sprintf("Message(id=%d) not found in Group(id=%d)", rl.MessageID, rl.GroupID)}
	}
	return nil
}

