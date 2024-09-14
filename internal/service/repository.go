package service

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type MessageRepository interface {
	CreateMessage(message RequestBody) (RequestBody, error)
	GetAllMessages() ([]RequestBody, error)
	UpdateMessageByID(id int, message RequestBody) (RequestBody, error)
	DeleteMessageByID(id int, message RequestBody) (RequestBody, error)
}

type MsgRepository struct {
	db *gorm.DB
}

func NewMessageRepository(db *gorm.DB) *MsgRepository {
	return &MsgRepository{db: db}
}

func (r *MsgRepository) CreateMessage(message RequestBody) (RequestBody, error) {
	result := r.db.Create(&message)
	if result.Error != nil {
		return RequestBody{}, result.Error
	}
	return message, nil
}

func (r *MsgRepository) GetAllMessages() ([]RequestBody, error) {
	var messages []RequestBody
	err := r.db.Find(&messages).Error
	if err != nil {
		return []RequestBody{}, err
	}
	return messages, err
}

func (r *MsgRepository) UpdateMessageByID(id int, message RequestBody) (RequestBody, error) {
	result := r.db.Clauses(clause.Returning{}).Where("id = ?", id).Updates(&message)
	if result.Error != nil {
		return message, result.Error
	}
	return message, nil
}

func (r *MsgRepository) DeleteMessageByID(id int, message RequestBody) (RequestBody, error) {
	result := r.db.Clauses(clause.Returning{}).Where("id = ?", id).Delete(&message)
	if result.Error != nil {
		return RequestBody{}, result.Error
	}
	return message, nil
}
