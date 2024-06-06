package model

import (
	"backend/pkg/error/externalerror"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type Status string
type Type string

const (
	StatusOngoing   Status = "ongoing"
	StatusAutoreply Status = "autoreply"
	StatusPending   Status = "pending"
	StatusClosed    Status = "closed"
	StatusReviewed  Status = "reviewed"
)

const (
	TypeUnknown Type = "unknown"
	TypeQuery   Type = "query"
	TypeRequest Type = "request"
)

type RequestQuery struct {
	gorm.Model
	Status    Status
	Type      Type
	BookingID *uint
	Booking   *Booking
	ChatID    uint
	Chat      Chat
	Messages  []Message
}

var ErrRequestHasNilBookingId = externalerror.AuthRequiredError{}
var ErrBookingIdDoesNotExist = errors.New("booking id does not exist")

func (r *RequestQuery) Create(db *gorm.DB) error {
	return db.Create(r).Error
}

func (r *RequestQuery) Update(db *gorm.DB) error {
	return db.Updates(r).Error
}

func (r *RequestQuery) Delete(db *gorm.DB) error {
	return db.Delete(r).Error
}

func (r *RequestQuery) BeforeSave(tx *gorm.DB) error {
	if r.Type == TypeRequest && r.BookingID == nil {
		return ErrRequestHasNilBookingId
	}
	if r.BookingID != nil {
		var booking Booking
		tx.First(&booking, *r.BookingID)
		if tx.Error != nil || booking.ID == 0 {
			return ErrBookingIdDoesNotExist
		}
	}
	return nil
}

func (m *RequestQuery) PopulateRequestQueries(db *gorm.DB) {
	var v1 uint = 1
	var v2 uint = 2
	var v3 uint = 3
	var v4 uint = 4

	requestQueries := []RequestQuery{
		{
			Status:    StatusClosed,
			Type:      TypeQuery,
			BookingID: &v1,
			ChatID:    1,
			Messages:  []Message{},
		},
		{
			Status:    StatusOngoing,
			Type:      TypeRequest,
			BookingID: &v2,
			ChatID:    1,
			Messages:  []Message{},
		},
		{
			Status:    StatusPending,
			Type:      TypeUnknown,
			BookingID: &v3,
			ChatID:    2,
			Messages:  []Message{},
		},
		{
			Status:    StatusReviewed,
			Type:      TypeQuery,
			BookingID: &v4,
			ChatID:    3,
			Messages:  []Message{},
		},
	}

	if err := db.Where("true").Unscoped().Delete(&RequestQuery{}).Error; err != nil {
		panic("failed to clear table")
	}

	for _, requestQuery := range requestQueries {
		err := db.Save(&requestQuery).Error
		if err != nil {
			fmt.Printf("Error when creating requestQuery")
		}
	}
}
