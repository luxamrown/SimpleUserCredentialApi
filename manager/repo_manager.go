package manager

import "enigmacamp.com/account/repository"

type RepoManager interface {
	AccountRepo() repository.AccountRepo
}

type repoManager struct {
	infra Infra
}

func (r *repoManager) AccountRepo() repository.AccountRepo {
	return repository.NewAccountRepo(r.infra.SqlDb())
}

func NewRepoManager(infra Infra) RepoManager {
	return &repoManager{
		infra: infra,
	}
}
