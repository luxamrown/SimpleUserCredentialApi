package usecase

import "enigmacamp.com/account/repository"

type RegisterUseCase interface {
	Register(userId, userEmail, userName, userPassword string)
}

type registerUseCase struct {
	repo repository.AccountRepo
}

func (r *registerUseCase) Register(userId, userEmail, userName, userPassword string) {
	r.repo.Register(userId, userEmail, userName, userPassword)
}

func NewRegisterUseCase(repo repository.AccountRepo) RegisterUseCase {
	return &registerUseCase{
		repo: repo,
	}
}
