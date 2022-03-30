package manager

import "enigmacamp.com/account/usecase"

type UseCaseManager interface {
	LoginUserUsecase() usecase.LoginUserUsecase
	RegisterUseCase() usecase.RegisterUseCase
}

type useCaseManager struct {
	repo RepoManager
}

func (u *useCaseManager) LoginUserUsecase() usecase.LoginUserUsecase {
	return usecase.NewLoginUseCase(u.repo.AccountRepo())
}

func (u *useCaseManager) RegisterUseCase() usecase.RegisterUseCase {
	return usecase.NewRegisterUseCase(u.repo.AccountRepo())
}

func NewUseCaseManager(manager RepoManager) UseCaseManager {
	return &useCaseManager{
		repo: manager,
	}
}
