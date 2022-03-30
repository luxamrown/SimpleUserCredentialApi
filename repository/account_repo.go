package repository

import "enigmacamp.com/account/model"

type AccountRepo interface {
	Login(user model.Account) error
	Register(userId string, userEmail string, userName string, userPassword string)
}
