package usecase

import (
	"enigmacamp.com/account/model"
	"enigmacamp.com/account/repository"
)

type LoginUserUsecase interface {
	Login(userName, userPassword string) error
}

type loginUserUsecase struct {
	repo repository.AccountRepo
}

func (l *loginUserUsecase) Login(userName, userPassword string) error {
	userAuth := model.NewAccount(userName, userPassword)
	err := l.repo.Login(userAuth)
	if err != nil {
		return err
	}
	return nil
}

func NewLoginUseCase(repo repository.AccountRepo) LoginUserUsecase {
	return &loginUserUsecase{
		repo: repo,
	}
}
