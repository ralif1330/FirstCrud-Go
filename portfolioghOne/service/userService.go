package service

import (
	"portfolioghOne/models"
	"portfolioghOne/repositories"

	"github.com/gin-gonic/gin"
)

type UsersUserService struct {
	UsersUserServiceRepo *repositories.UsersUserRepo
}

func NewUsersUserService(UsersUserRepo *repositories.UsersUserRepo) *UsersUserService {
	return &UsersUserService{
		UsersUserServiceRepo: UsersUserRepo,
	}
}

func (uus UsersUserService) ListUsersUserService(ctx *gin.Context) ([]*models.UsersUser, *models.ResponseError) {
	return uus.UsersUserServiceRepo.ListUsersUserRepo(ctx)
}
