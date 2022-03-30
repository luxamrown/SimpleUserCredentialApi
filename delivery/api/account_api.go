package api

import (
	"enigmacamp.com/account/delivery/appreq"
	"enigmacamp.com/account/manager"
	"github.com/gin-gonic/gin"
)

type AccountApi struct {
	BaseApi
	usecase manager.UseCaseManager
}

func (a *AccountApi) loginAccount() gin.HandlerFunc {
	return func(c *gin.Context) {
		var accReq appreq.AccountRequest
		err := a.ParseRequestBody(c, &accReq)
		if err != nil {
			a.FailedRequest(c, "api-loginAccount", "02", "Can not parse body")
			return
		}
		err = a.usecase.LoginUserUsecase().Login(accReq.UserName, accReq.UserPassword)
		if err != nil {
			a.Failed(c, "api-loginAccount", "01", "USER NOT FOUND..")
			return
		}
		a.Succes(c, "Login Account", accReq)
	}
}

func (a *AccountApi) registerAccount() gin.HandlerFunc {
	return func(c *gin.Context) {
		var accReq appreq.AccountRequestRegister
		err := a.ParseRequestBody(c, &accReq)
		if err != nil {
			a.FailedRequest(c, "api-regisAccount", "02", "Can not parse body")
			return
		}
		a.usecase.RegisterUseCase().Register(accReq.UserId, accReq.UserEmail, accReq.UserName, accReq.UserPassword)
		a.Succes(c, "Register Account", accReq)
	}
}

func NewAccountApi(accountRoute *gin.RouterGroup, usecase manager.UseCaseManager) {
	api := AccountApi{
		usecase: usecase,
	}
	accountRoute.POST("/login", api.loginAccount())
	accountRoute.POST("/register", api.registerAccount())
}
