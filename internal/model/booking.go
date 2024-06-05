package model

import (
	"gorm.io/gorm"
)

// Booking represents our "authentication" data.
// Our "authentication" process should be just adding to this table
// after trivially assuming they are all correct.
type Booking struct {
	gorm.Model
	RoomNumber string
	LastName   string
	ChatId     uint
	Chat       Chat
}

func (b *Booking) Create(db *gorm.DB) error {
	return db.Create(b).Error
}

func (b *Booking) Update(db *gorm.DB) error {
	return db.Updates(b).Error
}

func (b *Booking) Delete(db *gorm.DB) error {
	return db.Delete(b).Error
}

// populate bookings table with sample values

var sampleBookingWithRequest = Booking{
	RoomNumber: "Testing Room 1",
	LastName:   "Guest1",
	ChatId:     1,
}

var sampleBookingWithQuery = Booking{
	RoomNumber: "Testing Room 2",
	LastName:   "Guest2",
	ChatId:     2,
}

var sampleBookingWithUnknown = Booking{
	RoomNumber: "Testing Room 3",
	LastName:   "Guest3",
	ChatId:     3,
}

func (*Booking) PopulateBookings(db *gorm.DB) {
	if err := db.Where("true").Unscoped().Delete(&Booking{}).Error; err != nil {
		panic("failed to clear table")
	}
	db.Create(&sampleBookingWithRequest)
	db.Create(&sampleBookingWithQuery)
	db.Create(&sampleBookingWithUnknown)
}
