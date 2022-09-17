package models

import (
	"errors"
	"gorm.io/gorm"
	"golang.org/x/crypto/bcrypt"
	"time"
  )

type User struct {
	ID 			string  `gorm:"column:id;default:uuid_generate_v4()"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
	Name     	string  `gorm:"column:name"`
	Email       string  `gorm:"column:email;not null;"`
	Password    string  `gorm:"column:password;not null;"`
	Document    string  `gorm:"column:cpf;not null;"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	if len(u.Password) == 0 {
		return errors.New("password should not be empty")
	}
	bytePassword := []byte(u.Password)
	// Make sure the second param `bcrypt generator cost` between [4, 32)
	passwordHash, _ := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	u.Password = string(passwordHash)
	return nil
}

func (t *User) TableName() string {
    return "users"
}