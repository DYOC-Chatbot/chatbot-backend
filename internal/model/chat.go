package model

import (
	"gorm.io/gorm"
)

// Chat represents the telegram chat and has information required to send messages to it
// For simplicity, assume one booking to one chat for now. We can change this in future iterations if time permits.
type Chat struct {
	gorm.Model
	TelegramChatId int64 `gorm:"unique"`
	Booking        *Booking
	RequestQueries []RequestQuery
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

// populate chat table with sample values
var SampleChatWithRequest = Chat{
	TelegramChatId: 1,
	RequestQueries: []RequestQuery{},
}

var sampleChatWithQuery = Chat{
	TelegramChatId: 2,
	RequestQueries: []RequestQuery{},
}

var sampleChatWithUnknown = Chat{
	TelegramChatId: 3,
	RequestQueries: []RequestQuery{},
}

func (*Chat) PopulateChats(db *gorm.DB) {
	if err := db.Where("true").Unscoped().Delete(&Chat{}).Error; err != nil {
		panic("failed to clear table")
	}
	db.Create(&SampleChatWithRequest)
	db.Create(&sampleChatWithQuery)
	db.Create(&sampleChatWithUnknown)
}
