package model

import (
	"backend/internal/viewmodel"
	"errors"
	"fmt"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Role string

const (
	StaffRole Role = "staff"
	AdminRole Role = "admin"
)

type User struct {
	gorm.Model
	Username          string `gorm:"unique"`
	EncryptedPassword string
	Messages          []Message `gorm:"foreignKey:HotelStaffId"`
	Role              Role
}

func (u *User) ToView() *viewmodel.UserView {
	return &viewmodel.UserView{
		ID: u.ID, Username: u.Username,
	}
}

func (u *User) Create(db *gorm.DB) error {
	err := db.Create(u).Error
	if errors.Is(err, gorm.ErrDuplicatedKey) {
		return errors.New("username already taken")
	}
	return err
}

func (u *User) Update(db *gorm.DB) error {
	return db.Updates(u).Error
}

func (u *User) Delete(db *gorm.DB) error {
	return db.Delete(u).Error
}

func (u *User) BeforeCreate(db *gorm.DB) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(u.EncryptedPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.EncryptedPassword = string(bytes)
	return nil
}

func PopulateUsers(db *gorm.DB) {
	users := []User{
		{
			Username:          "staff1",
			EncryptedPassword: "unencryptedsamplepassword",
			Messages:          []Message{},
			Role:              StaffRole,
		},
		{
			Username:          "staff2",
			EncryptedPassword: "unencryptedsamplepassword",
			Messages:          []Message{},
			Role:              AdminRole,
		},
	}

	if err := db.Where("true").Unscoped().Delete(&User{}).Error; err != nil {
		panic("failed to clear table")
	}

	for _, user := range users {
		err := db.Save(&user).Error
		if err != nil {
			fmt.Printf("Error when creating chat")
		}
	}
}
