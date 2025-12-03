package controller

import (
	"autoplay-hub/dao/mysql"
	"autoplay-hub/logic"
	"autoplay-hub/models"
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func RegisterHandler(c *gin.Context) {
	// 1.获取参数以及参数校验
	p := new(models.ParamRegister)
	// 参数错误，处理错误并返回响应
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("RegisterHandler PostParam Error", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	// 2.业务处理
	if err := logic.Register(p); err != nil {
		zap.L().Error("logic.Register Error", zap.Error(err))
		if errors.Is(err, mysql.ErrorUserExist) {
			ResponseError(c, CodeUserExists)
			return
		} else {
			ResponseError(c, CodeServerBusy)
			return
		}
	}
	// 3. 返回响应
	ResponseSuccess(c, nil)
}

func LoginHandler(c *gin.Context) {
	// 1.参数校验
	p := new(models.ParamLogin)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("LoginHandler PostParam Error", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	// 2.业务处理
	user, err := logic.Login(p)
	if err != nil {
		zap.L().Error("logic.Login error", zap.Error(err))
		if errors.Is(err, mysql.ErrorUserExist) {
			ResponseError(c, CodeUserExists)
			return
		} else if errors.Is(err, mysql.ErrorInvalidPassword) {
			ResponseError(c, CodeInvalidPassword)
			return
		} else {
			ResponseError(c, CodeServerBusy)
			return
		}
	}
	// 3.返回响应
	ResponseSuccess(c, gin.H{
		"user_id":   fmt.Sprintf("%d", user.UserID),
		"user_name": user.Username,
		"token":     user.Token,
		"is_admin":  user.IsAdmin,
	})
}
