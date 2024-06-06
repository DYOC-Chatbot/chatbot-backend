package model

import (
	"fmt"

	"gorm.io/gorm"
)

// Chat represents the telegram chat and has information required to send messages to it
// For simplicity, assume one booking to one chat for now. We can change this in future iterations if time permits.
type Chat struct {
	gorm.Model
	TelegramChatId int64 `gorm:"unique"`

	Booking        *Booking       `gorm:"->"`
	RequestQueries []RequestQuery `gorm:"->;<-"`
}

func (c *Chat) Create(db *gorm.DB) error {
	return db.Create(c).Error
}

func (c *Chat) Update(db *gorm.DB) error {
	return db.Updates(c).Error
}

func (c *Chat) Delete(db *gorm.DB) error {
	return db.Delete(c).Error
}

<<<<<<< HEAD
func PopulateChats(db *gorm.DB) {
=======
func (*Chat) PopulateChats(db *gorm.DB) {
>>>>>>> ee114e6 ((refactor:) decouple seeding from server)
	chats := []Chat{
		{
			TelegramChatId: 1,
			RequestQueries: []RequestQuery{},
		},
		{
			TelegramChatId: 2,
			RequestQueries: []RequestQuery{},
		},
		{
			TelegramChatId: 3,
			RequestQueries: []RequestQuery{},
		},
		{
			TelegramChatId: 4,
			RequestQueries: []RequestQuery{},
		},
	}

<<<<<<< HEAD
	for _, chat := range chats {
		err := db.Save(&chat).Error
		if err != nil {
			fmt.Printf("Error when creating chat")
		}
	}
=======
	if err := db.Where("true").Unscoped().Delete(&Chat{}).Error; err != nil {
		panic("failed to clear table")
	}

	for _, chat := range chats {
		err := db.Save(&chat).Error
		if err != nil {
			fmt.Printf("Error when creating chat")
		}
	}
>>>>>>> ee114e6 ((refactor:) decouple seeding from server)
}
