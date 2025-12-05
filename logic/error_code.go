package logic

import "errors"

var (
	ErrorInvalidUserID = errors.New("用户不对应")
	ErrorNoDevices     = errors.New("没有可用的设备")
)
