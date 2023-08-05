package repositories

import (
	"database/sql"
	"net/http"
	"portfolioghOne/models"
	"portfolioghOne/repositories/dbContext"

	"github.com/gin-gonic/gin"
)

type UsersUserRepo struct {
	DbHandler   *sql.DB
	Transaction *sql.Tx
}

func NewUserUserRepository(dbHandler *sql.DB) *UsersUserRepo {
	return &UsersUserRepo{
		DbHandler: dbHandler,
	}
}

func (uur UsersUserRepo) ListUsersUserRepo(ctx *gin.Context) ([]*models.UsersUser, *models.ResponseError) {

	store := dbContext.New(uur.DbHandler)
	usersUserList, err := store.ListUsersUserImpl(ctx)

	listUsersUser := make([]*models.UsersUser, 0)

	for _, v := range usersUserList {
		usersUser := &models.UsersUser{
			UserEntityID: v.UserEntityID,
			UserName:     v.UserName,
			UserEmail:    v.UserEmail,
			UserGender:   v.UserGender,
		}
		listUsersUser = append(listUsersUser, usersUser)
	}

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}
	return listUsersUser, nil
}
