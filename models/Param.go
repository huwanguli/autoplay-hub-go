package models

import "time"

// ParamRegister RegisterPOST请求中使用
type ParamRegister struct {
	Username   string `json:"username" binding:"required"`
	Password   string `json:"password" binding:"required"`
	RePassword string `json:"re_password" binding:"required,eqfield=Password"`
	IsAdmin    bool   `json:"is_admin" `
}

type ParamLogin struct {
	Username string `json:"username" binding:"required" binding:"required"`
	Password string `json:"password" binding:"required" binding:"required"`
}

// ParamScript 创建脚本POST中使用
type ParamScript struct {
	ScriptName  string `json:"name" binding:"required"`
	Description string `json:"description"`
	OwnerID     int64  `json:"owner_id"`
	Content     string `json:"content" binding:"required"`
}

// ParamUpdateScript 修改脚本内容用结构体
type ParamUpdateScript struct {
	ScriptName  *string `json:"name"`
	Description *string `json:"description"`
	Content     *string `json:"content"`
}

type ParamScriptRun struct {
	ScriptID  int64  `json:"script_id,string"`
	DeviceUrl string `json:"device_url" binding:"required"`
}

type ParamStopTask struct {
	TaskID     int64     `json:"task_id,string"`
	Status     int64     `json:"status" binding:"required"`
	ExecutedAt time.Time `json:"executed_at"`
}
