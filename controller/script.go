package controller

import (
	"autoplay-hub/dao/mysql"
	"autoplay-hub/logic"
	"autoplay-hub/models"
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// ScriptHandler 创建脚本
func ScriptHandler(c *gin.Context) {
	// 1.参数的校验
	p := new(models.ParamScript)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("ScriptHandler param err", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	userID, err := getCurrentUserID(c)
	if err != nil {
		zap.L().Error("ScriptHandler getCurrentUserID err", zap.Error(err))
		ResponseError(c, CodeNeedLogin)
		return
	}
	p.OwnerID = userID
	// 2.业务处理
	if err := logic.CreateScript(p); err != nil {
		zap.L().Error("ScriptHandler logic.CreateScript err", zap.Error(err))
		if errors.Is(err, mysql.ErrorScriptExist) {
			ResponseError(c, CodeScriptExists)
			return
		} else if err != nil {
			ResponseError(c, CodeServerBusy)
			return
		}
	}
	// 3.返回响应
	ResponseSuccess(c, nil)
}

// AllScriptInfoHandler 获取所有脚本
func AllScriptInfoHandler(c *gin.Context) {
	// 1. 参数的校验 (查询参数，页数以及每一页的容量和排序依据)
	page, size := getPageInfo(c)
	// 获取参数
	permission, err := getCurrentUserIsAdmin(c)
	if errors.Is(err, ErrorUserNotLogin) {
		ResponseError(c, CodeNeedLogin)
		return
	} else if err != nil {
		ResponseError(c, CodeServerBusy)
		return
	}
	userID, err := getCurrentUserID(c)
	if errors.Is(err, ErrorUserNotLogin) {
		ResponseError(c, CodeNeedLogin)
		return
	} else if err != nil {
		ResponseError(c, CodeServerBusy)
		return
	}
	if permission == true {
		userID = 0
	}
	data, err := logic.GetAllScriptList(page, size, userID)
	if err != nil {
		zap.L().Error("logic.GetAllScriptList err", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)
}

// ScriptDetailHandler 获取脚本详情
func ScriptDetailHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		ResponseError(c, CodeInvalidParam)
		return
	}
	userID, err := getCurrentUserID(c)
	if errors.Is(err, ErrorUserNotLogin) {
		ResponseError(c, CodeNeedLogin)
		return
	} else if err != nil {
		ResponseError(c, CodeServerBusy)
		return
	}
	data, err := logic.GetScriptDetail(id, userID)
	if err != nil {
		if errors.Is(err, mysql.ErrorScriptNotExist) {
			ResponseError(c, CodeScriptNotExists)
			return
		} else if errors.Is(err, logic.ErrorInvalidUserID) {
			ResponseError(c, CodeInvalidUser)
			return
		}
		zap.L().Error("logic.GetScriptDetail err", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)
}

// UpdateScriptHandler 编辑脚本功能实现
func UpdateScriptHandler(c *gin.Context) {
	// 1. 参数的校验
	p := new(models.ParamUpdateScript)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("UpdateScriptHandler param err", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	// 2.业务的处理
	idStr := c.Param("id")
	if idStr == "" {
		ResponseError(c, CodeInvalidParam)
		return
	}
	id, _ := strconv.ParseInt(idStr, 10, 64)
	ownerID, err := getCurrentUserID(c)
	if errors.Is(err, ErrorUserNotLogin) {
		ResponseError(c, CodeNeedLogin)
		return
	}
	if err := logic.UpdateScript(p, id, ownerID); err != nil {
		zap.L().Error("UpdateScriptHandler logic.UpdateScript err", zap.Error(err))
		if errors.Is(err, mysql.ErrorScriptNotExist) {
			ResponseError(c, CodeScriptNotExists)
			return
		}
		ResponseError(c, CodeServerBusy)
		return
	}
	// 3.返回响应
	ResponseSuccess(c, nil)
}
