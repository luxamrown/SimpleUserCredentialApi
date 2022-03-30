package repository

import (
	"errors"

	"enigmacamp.com/account/model"
	"github.com/jmoiron/sqlx"
)

type AccountRepoImpl struct {
	AccountDb *sqlx.DB
}

func (a *AccountRepoImpl) Login(user model.Account) error {
	var isUserExists int
	err := a.AccountDb.Get(&isUserExists, "SELECT COUNT(id) FROM account_list WHERE user_name = $1 and user_password = $2 and is_blocked = $3", user.UserName, user.UserPasword, "false")
	if err != nil {
		return err
	}
	if isUserExists == 0 {
		return errors.New("USERRRRRR NOT FOUND")
	}
	return nil
}

func (a *AccountRepoImpl) Register(userId string, userEmail string, userName string, userPassword string) {
	tx := a.AccountDb.MustBegin()
	tx.MustExec("INSERT INTO account_list(id, user_email, user_name, user_password, is_blocked) VALUES($1, $2, $3, $4, $5)", userId, userEmail, userName, userPassword, "false")
	tx.Commit()
}

func NewAccountRepo(AccountDb *sqlx.DB) AccountRepo {
	accountRepo := AccountRepoImpl{
		AccountDb: AccountDb,
	}
	return &accountRepo
}
