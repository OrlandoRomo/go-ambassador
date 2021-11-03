package model

import (
	"golang.org/x/crypto/bcrypt"
)

type Admin User
type Ambassador User

type User struct {
	ID           uint    `json:"user_id"`
	FirstName    string  `json:"first_name"`
	LastName     string  `json:"last_name"`
	Email        string  `json:"email" gorm:"unique"`
	Password     []byte  `json:"-"`
	IsAmbassador bool    `json:"-"`
	Revenue      float64 `json:"revenue,omitempty" gorm:"-"`
}

func (u *User) SetPassword(password string) {
	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 12)
	u.Password = hashPassword
}

func (u *User) ComparePassword(password string) error {
	return bcrypt.CompareHashAndPassword(u.Password, []byte(password))
}
