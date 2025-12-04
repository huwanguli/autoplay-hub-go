package controller

import "C"
import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
)

const (
	CtxUserIDKey  = "userID"
	CtxIsAdminKey = "isAdmin"
)

var (
	ErrorUserNotLogin = errors.New("用户未登录")
	ErrorInvalidID    = errors.New("无效的ID")
)

// 获取当前登录用户ID
func getCurrentUserID(c *gin.Context) (userID int64, err error) {
	uid, ok := c.Get(CtxUserIDKey)
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	userID, ok = uid.(int64)
	if !ok {
		err = ErrorInvalidID
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
		err = ErrorInvalidID
		return
	}
	return
}

// 获取页码配置
func getPageInfo(c *gin.Context) (int64, int64) {
	pageStr := c.Query("page")
	sizeStr := c.Query("size")
	var (
		page int64
		size int64
		err  error
	)
	page, err = strconv.ParseInt(pageStr, 10, 64)
	if err != nil {
		page = 1
	}
	size, err = strconv.ParseInt(sizeStr, 10, 64)
	if err != nil {
		size = 10
	}
	return page, size
}
