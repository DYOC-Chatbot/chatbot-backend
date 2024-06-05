package model

import (
	"time"

	"gorm.io/gorm"
)

// By is the type of entity that sent the message
type By string

const (
	ByGuest By = "guest"
	ByBot   By = "bot"
	ByStaff By = "staff"
)

type Message struct {
	gorm.Model
	TelegramMessageId int64
	By                By
	MessageBody       string
	Timestamp         time.Time
	HotelStaffId      *uint
	HotelStaff        *User
	RequestQueryId    uint
	RequestQuery      RequestQuery
}

func (m *Message) Create(db *gorm.DB) error {
	return db.Create(m).Error
}

func (m *Message) Update(db *gorm.DB) error {
	return db.Updates(m).Error
}

func (m *Message) Delete(db *gorm.DB) error {
	return db.Delete(m).Error
}

var v uint = 1

var sampleGuestMessage = Message{
	TelegramMessageId: 1,
	By:                ByGuest,
	MessageBody:       "Hello",
	Timestamp:         time.Now(),
	HotelStaffId:      &v,
	RequestQueryId:    1,
}

func (m *Message) PopulateMessages(db *gorm.DB) {
	if err := db.Where("true").Unscoped().Delete(&Message{}).Error; err != nil {
		panic("failed to clear table")
	}
	db.Create(&sampleGuestMessage)
}
