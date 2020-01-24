// Copyright (c) 2020 Vorotynsky Maxim

package data

import "microAuth/model"

type UserRepository interface {
	GetUserById(id int32) (model.User, error)
	GetUserByName(userName string) (model.User, error)
	Create(user *model.User) error
	Update(user *model.User) error
	DeleteById(id int32) error
	DeleteByName(userName string) error
}
