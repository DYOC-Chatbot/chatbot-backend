package chat

import (
	"backend/internal/dataaccess/requestquery"
	"backend/internal/model"

	"gorm.io/gorm"
)

const invalidID = "invalid chat id"

func Read(db *gorm.DB, id uint) (*model.Chat, error) {
	var chat model.Chat
	result := db.Model(&model.Chat{}).
		Where("id = ?", id).
		First(&chat, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &chat, nil
}

func ReadByTgChatID(db *gorm.DB, tgChatID int64) (*model.Chat, error) {
	var chat model.Chat
	result := db.Model(&model.Chat{}).
		Where("telegram_chat_id = ?", tgChatID).
		First(&chat)
	if result.Error != nil {
		return nil, result.Error
	}

	return &chat, nil
}

func Create(db *gorm.DB, chat *model.Chat) error {
	return chat.Create(db)
}

func Update(db *gorm.DB, chat *model.Chat) error {
	return chat.Update(db)
}

func Delete(db *gorm.DB, chat *model.Chat) error {
	err := db.Transaction(func(tx *gorm.DB) error {
		if err := pruneAssociations(tx, chat); err != nil {
			return err
		}

		return chat.Delete(tx)
	})

	return err
}

func pruneAssociations(db *gorm.DB, chat *model.Chat) error {
	var rqqs []model.RequestQuery
	result := db.Model(&model.RequestQuery{}).
		Where("chat_id = ?", chat.ID).
		Find(&rqqs)
	if result.Error != nil {
		return result.Error
	}

	for _, rqq := range rqqs {
		if err := requestquery.Delete(db, &rqq); err != nil {
			return err
		}
	}

	return nil
}
