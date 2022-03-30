package model

type Account struct {
	UserName    string `db:"user_name"`
	UserPasword string `db:"user_password"`
}

func (a *Account) GetUserName() string {
	return a.UserName
}

func (a *Account) GetUserPassword() string {
	return a.UserPasword
}

func (a *Account) SetUserName(userName string) {
	a.UserName = userName
}

func (a *Account) SetUserPassword(userPassword string) {
	a.UserPasword = userPassword
}

func NewAccount(userName, userPassword string) Account {
	return Account{
		UserName:    userName,
		UserPasword: userPassword,
	}
}
