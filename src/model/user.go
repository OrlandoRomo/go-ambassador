package model

import "golang.org/x/crypto/bcrypt"

type User struct {
	ID           uint   `json:"user_id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Email        string `json:"email" gorm:"unique"`
	Password     []byte `json:"-"`
	IsAmbassador bool   `json:"-"`
}

func NewUser(data map[string]string) *User {
	return &User{
		FirstName: data["first_name"],
		LastName:  data["last_name"],
		Email:     data["email"],
	}
}

func (u *User) SetPassword(password string) {
	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 12)
	u.Password = hashPassword
}

func (u *User) ComparePassword(password string) error {
	return bcrypt.CompareHashAndPassword(u.Password, []byte(password))
}
