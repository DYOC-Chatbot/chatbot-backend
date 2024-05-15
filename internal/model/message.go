package model

import (
	"gorm.io/gorm"
	"time"
)

// By is the type of entity that sent the message
type By string

const (
	Guest By = "guest"
	Bot   By = "bot"
	Staff By = "staff"
)

type Message struct {
	gorm.Model
	TelegramMessageId int64
	By                By
	MessageBody       string
	Timestamp         time.Time
	HotelStaffId      uint
	HotelStaff        User
	RequestQueryId    uint
	RequestQuery      User
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