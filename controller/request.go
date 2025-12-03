package controller

import "C"
import (
	"errors"

	"github.com/gin-gonic/gin"
)

const (
	CtxUserIDKey  = "userID"
	CtxIsAdminKey = "isAdmin"
)

var ErrorUserNotLogin = errors.New("用户未登录")

// 获取当前登录用户ID
func getCurrentUserID(c *gin.Context) (userID int64, err error) {
	uid, ok := c.Get(CtxUserIDKey)
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	userID, ok = uid.(int64)
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	return
}

// 获取当前登录用户权限
func getCurrentUserIsAdmin(c *gin.Context) (isAdmin bool, err error) {
	isadmin, ok := c.Get(CtxIsAdminKey)
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	isAdmin, ok = isadmin.(bool)
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	return
}
