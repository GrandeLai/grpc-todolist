package handler

import (
	"api-gateway/internal/service"
	"api-gateway/pkg/e"
	"api-gateway/pkg/resp"
	"api-gateway/pkg/util"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UserRegisterHandler(c *gin.Context) {
	var userReq service.UserRequest
	PanicIfUserError(c.Bind(&userReq))
	//gin.Key中取出服务实例
	userService := c.Keys["user"].(service.UserServiceClient)
	userResp, err := userService.UserRegister(context.Background(), &userReq)
	PanicIfUserError(err)
	r := resp.Response{
		Data:   userResp,
		Msg:    e.GetMsg(uint(userResp.Code)),
		Status: uint(userResp.Code),
	}
	c.JSON(http.StatusOK, r)
}

func UserLoginHandler(c *gin.Context) {
	var userReq service.UserRequest
	PanicIfUserError(c.Bind(&userReq))
	//gin.Key中取出服务实例
	userService := c.Keys["user"].(service.UserServiceClient)
	userResp, err := userService.UserLogin(context.Background(), &userReq)
	PanicIfUserError(err)
	token, err := util.GenerateToken(uint(userResp.UserDetail.UserID))
	r := resp.Response{
		Data: resp.TokenData{
			User:  userResp,
			Token: token,
		},
		Msg:    e.GetMsg(uint(userResp.Code)),
		Status: uint(userResp.Code),
	}
	c.JSON(http.StatusOK, r)
}
