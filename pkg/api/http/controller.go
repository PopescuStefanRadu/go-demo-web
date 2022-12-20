package http

import (
	"faceit/pkg/errors"
	"faceit/pkg/user"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserController struct {
	Service user.Service
}

func (c UserController) GetUsersByQuery(ctx *gin.Context) {
	query := user.FilterQuery{}
	err := ctx.BindQuery(&query)
	if err != nil {
		_ = ctx.Error(err)
		ctx.Abort()
		// TODO handle in middleware
		return
	}
	users, err := c.Service.FindUsers(ctx, &query)
	if err != nil {
		_ = ctx.Error(err)
		ctx.Abort()
		ctx.JSON(http.StatusInternalServerError, ErrorResponse{GlobalErrors: []string{err.Error()}})
		return
	}
	if len(users.Data) == 0 {
		err := errors.LibErr{
			Fault: errors.FaultClient,
			Err:   nil,
			Msg:   "could not find any users that match your query",
		}
		_ = ctx.Error(err)
		ctx.JSON(http.StatusNotFound, ErrorResponse{GlobalErrors: []string{err.Error()}})
		return
	}
	ctx.JSON(http.StatusOK, users)
}

func (c UserController) CreateUser(ctx *gin.Context) {
	usr := user.User{}
	err := ctx.Bind(&usr)
	if err != nil {
		_ = ctx.Error(err)
		ctx.Abort()
		// TODO err handling
		return
	}
	usr, err = c.Service.AddUser(ctx, usr)
	if err != nil {
		_ = ctx.Error(err)
		ctx.Abort()
		ctx.JSON(http.StatusInternalServerError, ErrorResponse{GlobalErrors: []string{err.Error()}})
		return
	}
	ctx.JSON(http.StatusOK, usr)
}

func (c UserController) UpdateUser(ctx *gin.Context) {
	update := user.UpdateUserInput{}
	err := ctx.Bind(update)
	if err != nil {
		_ = ctx.Error(err)
		ctx.Abort()
		return
	}
	updateUser, err := c.Service.UpdateUser(ctx, &update)
	if err != nil {
		_ = ctx.Error(err)
		ctx.Abort()
		ctx.JSON(http.StatusInternalServerError, ErrorResponse{GlobalErrors: []string{err.Error()}})
		return
	}
	ctx.JSON(http.StatusOK, updateUser)
}

func (c UserController) DeleteUser(ctx *gin.Context) {
	userId := ctx.Param("user_id")
	deletedUser, err := c.Service.DeleteUser(ctx, userId)
	if err != nil {
		_ = ctx.Error(err)
		ctx.Abort()
		ctx.JSON(http.StatusInternalServerError, ErrorResponse{GlobalErrors: []string{err.Error()}})
		return
	}
	ctx.JSON(http.StatusOK, deletedUser)
}
