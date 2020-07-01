package user

import (
	"github.com/jinzhu/gorm"
)

// Repository represent the users's repository contract
type Repository interface {
	GetByEmail(email string) (*User, error)
	Create(user *User) error
}

type userGorm struct {
	DB *gorm.DB
}

// NewRepository create an object that represent the Repository interface
func NewRepository(DB *gorm.DB) Repository {
	return &userGorm{DB}
}

func (ug *userGorm) Create(user *User) error {
	return ug.DB.Create(user).Error
}

func (ug *userGorm) GetByEmail(email string) (*User, error) {
	var u User
	if err := ug.DB.Where("email = ?", email).First(&u).Error; err != nil {
		return nil, err
	}
	return &u, nil
}
