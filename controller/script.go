package controller

import (
	"autoplay-hub/dao/mysql"
	"autoplay-hub/logic"
	"autoplay-hub/models"
	"errors"

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
		} else {
			ResponseError(c, CodeServerBusy)
			return
		}
	}
	// 3.返回响应
	ResponseSuccess(c, nil)
}
