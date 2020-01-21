// Copyright (c) 2020 Vorotynsky Maxim

package data

import "microAuth/model"

type UserRepository interface {
	GetUserById(id int32) *model.User
	GetUserByName(userName string) *model.User
	Create(user *model.User) error
	Update(user *model.User) error
	DeleteById(id int32) error
	DeleteByName(userName string) error
}
