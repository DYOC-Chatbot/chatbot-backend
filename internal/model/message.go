package model

import (
	"fmt"
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
	TelegramMessageID int64
	By                By
	MessageBody       string
	Timestamp         time.Time
	HotelStaffID      *uint
	RequestQueryID    uint

	RequestQuery *RequestQuery `gorm:"->"`
	HotelStaff   *User         `gorm:"->"`
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

func PopulateMessages(db *gorm.DB) {
	var v1 uint = 1

	messages := []Message{
		{
			TelegramMessageID: 1,
			By:                ByGuest,
			MessageBody:       "Hello",
			Timestamp:         time.Now(),
			RequestQueryID:    1,
		},
		{
			TelegramMessageID: 1,
			By:                ByBot,
			MessageBody:       "How may I help you",
			Timestamp:         time.Now(),
			RequestQueryID:    1,
		},
		{
			TelegramMessageID: 1,
			By:                ByGuest,
			MessageBody:       "Bye",
			Timestamp:         time.Now(),
			RequestQueryID:    1,
		},
		{
			TelegramMessageID: 1,
			By:                ByGuest,
			MessageBody:       "I would like extra pillows",
			Timestamp:         time.Now(),
			RequestQueryID:    2,
		},
		{
			TelegramMessageID: 1,
			By:                ByBot,
			MessageBody:       "Processing request",
			Timestamp:         time.Now(),
			RequestQueryID:    2,
		},
		{
			TelegramMessageID: 1,
			By:                ByStaff,
			MessageBody:       "Sending pillows to your room now",
			Timestamp:         time.Now(),
			HotelStaffID:      &v1,
			RequestQueryID:    2,
		},
		{
			TelegramMessageID: 2,
			By:                ByGuest,
			MessageBody:       "Food recommendation nearby",
			Timestamp:         time.Now(),
			RequestQueryID:    3,
		},
		{
			TelegramMessageID: 3,
			By:                ByGuest,
			MessageBody:       "Breakfast hours",
			Timestamp:         time.Now(),
			RequestQueryID:    4,
		},
		{
			TelegramMessageID: 3,
			By:                ByBot,
			MessageBody:       "Breakfast is at 0700 - 1000",
			Timestamp:         time.Now(),
			RequestQueryID:    4,
		},
	}

	for _, message := range messages {
		err := db.Save(&message).Error
		if err != nil {
			fmt.Printf("Error when creating message")
		}
	}
}
