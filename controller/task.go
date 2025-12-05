package controller

import (
	"autoplay-hub/dao/mysql"
	"autoplay-hub/logic"
	"autoplay-hub/models"
	"errors"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// ScriptRunHandler 脚本运行处理
func ScriptRunHandler(c *gin.Context) {
	// 1.参数的校验
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		zap.L().Error("GetTaskStop InvalidParam", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	p := new(models.ParamScriptRun)
	p.ScriptID = id
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("ParamScriptRun ShouldBindJSON error", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	// 2. 业务的处理
	// 获取当前用户ID
	userID, err := getCurrentUserID(c)
	if err != nil {
		zap.L().Error("getCurrentUserID error", zap.Error(err))
		ResponseError(c, CodeNeedLogin)
		return
	}
	// 插入初始化的任务消息
	if err := logic.ScriptRunFirst(p, userID); err != nil {
		zap.L().Error("logic.ScriptRun error", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	if err := logic.ScriptRun(p); err != nil {
		zap.L().Error("redis.ScriptRun error", zap.Error(err))
		ResponseError(c, CodeRunScriptFailed)
		return
	}
	// 建立websocket连接 (暂时搁置)
	// 3.返回响应
	ResponseSuccess(c, nil)
}

// TaskListHandler 获取脚本列表
func TaskListHandler(c *gin.Context) {
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
	data, err := logic.GetAllTaskList(page, size, userID)
	if err != nil {
		zap.L().Error("logic.GetAllScriptList err", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)
}

// TaskDetailHandler 获取任务详情
func TaskDetailHandler(c *gin.Context) {
	// 获取任务id
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
	data, err := logic.GetTaskDetail(id, userID)
	if err != nil {
		if errors.Is(err, mysql.ErrorScriptNotExist) {
			ResponseError(c, CodeScriptNotExists)
			return
		} else if errors.Is(err, logic.ErrorInvalidUserID) {
			ResponseError(c, CodeInvalidUser)
			return
		}
		zap.L().Error("logic.GetTaskDetail err", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)
}

// TaskStopHandler 暂停正在运行的任务
func TaskStopHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		zap.L().Error("GetTaskStop InvalidParam", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	p := new(models.ParamStopTask)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("ParamStop ShouldBindJSON error", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	p.TaskID = id
	p.ExecutedAt = time.Now()
	// 业务的处理
	if err := logic.TaskStop(p); err != nil {
		zap.L().Error("logic.TaskStop err", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	// 返回响应
	ResponseSuccess(c, nil)
}

// UpdateTaskHandler 更改任务详情
func UpdateTaskHandler(c *gin.Context) {
	// 1.参数的校验
	id, err := getCurrentRouteID(c)
	if err != nil {
		zap.L().Error("UpdateTaskHandler GetCurrentRouteID error InvalidParam", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	p := new(models.ParamUpdateTask)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("ParamUpdateTask ShouldBindJSON error", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	p.TaskID = id
	// 2. 业务的处理
	if err := logic.UpdateTask(p); err != nil {
		zap.L().Error("logic.UpdateTask err", zap.Error(err))
		if errors.Is(err, mysql.ErrorTaskNotExist) {
			ResponseError(c, CodeTaskNotExists)
			return
		}
		ResponseError(c, CodeServerBusy)
		return
	}
	// 3.返回响应
	ResponseSuccess(c, nil)
}

// DeleteTaskHandler 删除任务
func DeleteTaskHandler(c *gin.Context) {
	id, err := getCurrentRouteID(c)
	if err != nil {
		zap.L().Error("DeleteTaskHandler GetCurrentRouteID error InvalidParam", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	// 业务的处理
	if err := logic.DeleteTask(id); err != nil {
		zap.L().Error("logic.DeleteTask err", zap.Error(err))
		if errors.Is(err, mysql.ErrorTaskNotExist) {
			ResponseError(c, CodeTaskNotExists)
			return
		}
		ResponseError(c, CodeServerBusy)
		return
	}
	// 返回响应
	ResponseSuccess(c, nil)
}
