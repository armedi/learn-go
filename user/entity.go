package user

import (
	"github.com/jinzhu/gorm"
)

// User ...
type User struct {
	gorm.Model
	Email        string `gorm:"type:varchar(100);not null;unique_index"`
	Password     string `gorm:"-"`
	PasswordHash string `gorm:"column:password;type:varchar(100);not null"`
}

type authenticationRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

// RegisterRequest - data received in register request body
type RegisterRequest struct {
	authenticationRequest
}

// LoginRequest - data received in login request body
type LoginRequest struct {
	authenticationRequest
}

// LoginResponse - data sent in login response body
type LoginResponse struct {
	AccessToken string `json:"access_token"`
}
