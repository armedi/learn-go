package user

import "github.com/jinzhu/gorm"

// User ...
type User struct {
	gorm.Model
	Email    string `gorm:"type:varchar(100);not null;unique_index"`
	Password string `gorm:"type:varchar(100);not null"`
}

// RegisterRequest - data sent in request body
type RegisterRequest struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}
