package user

import (
	"backend/internal/model"
	"gorm.io/gorm"
)

func ReadAll(db *gorm.DB) []*model.User {
	var users []*model.User
	db.Find(&users)
	return users
}
