package mysql

import "errors"

var (
	ErrorUserExist       = errors.New("用户已存在")
	ErrorUserNotExist    = errors.New("用户不存在")
	ErrorInvalidPassword = errors.New("密码错误")
	ErrorInvalidID       = errors.New("无效的ID")
	ErrorScriptExist     = errors.New("脚本已存在")
	ErrorScriptNotExist  = errors.New("脚本不存在")
	ErrorUpdateFailed    = errors.New("更新失败")
	ErrorTaskNotExist    = errors.New("任务不存在")
)
