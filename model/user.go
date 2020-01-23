// Copyright (c) 2020 Vorotynsky Maxim

package model

import (
	"errors"
	"log"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	UserData     UserData
	PasswordHash []byte
}
type UserWithPurePassword struct {
	UserData UserData
	Password string
}
type UserData struct {
	UserID   int32
	UserName string
}

func NewUser(id int32, name string, password string) *User {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		log.Println("[NewUser]:", err)
		return nil
	}
	return &User{UserData{id, name}, hash}
}

func (usr *User) ConfirmPassword(password string) error {
	if usr == nil {
		return errors.New("User is null")
	}
	return bcrypt.CompareHashAndPassword(usr.PasswordHash, []byte(password))
}

func (usr *UserWithPurePassword) HashPassword() *User {
	hash, err := bcrypt.GenerateFromPassword([]byte(usr.Password), bcrypt.MinCost)
	if err != nil {
		log.Println("[NewUser]:", err)
		return nil
	}
	return &User{usr.UserData, hash}
}
