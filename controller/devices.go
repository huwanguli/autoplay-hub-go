package controller

import (
	"autoplay-hub/logic"
	"errors"
	"os/exec"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func GetDevicesHandler(c *gin.Context) {
	cmd := exec.Command("adb", "devices")
	output, err := cmd.Output()

	if err != nil {
		zap.L().Error("cmd 执行错误", zap.Any("cmd", cmd), zap.Error(err))
		ResponseError(c, CodeCmdRunFailed)
		return
	}
	// 解析输出
	devices, err := logic.GetDevices(output)
	if err != nil {
		if errors.Is(err, logic.ErrorNoDevices) {
			ResponseError(c, CodeNoDevices)
			return
		}
		ResponseError(c, CodeServerBusy)
		return
	}
	// 返回响应
	ResponseSuccess(c, devices)
}
