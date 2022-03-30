package appreq

type AccountRequest struct {
	UserName     string `json:"user_name"`
	UserPassword string `json:"user_password"`
}

type AccountRequestRegister struct {
	UserId       string `json:"id"`
	UserEmail    string `json:"user_email"`
	UserName     string `json:"user_name"`
	UserPassword string `json:"user_password"`
}
