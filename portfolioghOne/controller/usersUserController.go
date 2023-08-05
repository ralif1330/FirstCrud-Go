package controller

import (
	"net/http"
	"portfolioghOne/service"

	"github.com/gin-gonic/gin"
)

type UsersUserController struct {
	UsersUserService *service.UsersUserService
}

// DECLARE CONSTRUCTOR

func NewUsersUserController(UsersUserService *service.UsersUserService) *UsersUserController {
	return &UsersUserController{
		UsersUserService: UsersUserService,
	}
}

func (UsersUserController UsersUserController) ListUsersUserHttp(ctx *gin.Context) {
	response, responseErr := UsersUserController.UsersUserService.ListUsersUserService(ctx)

	if responseErr != nil {
		ctx.JSON(responseErr.Status, responseErr)
		return
	}

	ctx.JSON(http.StatusOK, response)
}
