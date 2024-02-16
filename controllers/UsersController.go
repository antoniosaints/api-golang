package controllers

import (
	obj "api/entities"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UsersController struct {
	users []obj.User
}

func UserControllers() *UsersController {
	return &UsersController{}
}

func (u *UsersController) GetUsers(ctx *gin.Context) {
	u.users = []obj.User{}
	u.users = append(u.users, obj.NewUser("João", 20, "2022-01-01"), obj.NewUser("Maria", 30, "2022-02-02"))
	ctx.JSON(http.StatusOK, u.users)

}

func (u *UsersController) GetUser(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, obj.GetUser(ctx.Param("name")))
}
