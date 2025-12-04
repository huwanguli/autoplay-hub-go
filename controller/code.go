package controller

//定义可能报的错

type ResCode int64

const (
	CodeSuccess ResCode = 1000 + iota
	CodeInvalidParam
	CodeUserExists
	CodeUserNotExists
	CodeInvalidPassword
	CodeServerBusy
	CodeNeedLogin
	CodeInvalidToken
	CodeScriptExists
	CodeScriptNotExists
	CodeInvalidUser
)

var codeMsgMap = map[ResCode]string{
	CodeSuccess:         "success",
	CodeInvalidParam:    "请求参数错误",
	CodeUserExists:      "用户不存在",
	CodeUserNotExists:   "用户已存在",
	CodeInvalidPassword: "用户名或密码错误",
	CodeServerBusy:      "服务器繁忙",
	CodeNeedLogin:       "需要登录",
	CodeInvalidToken:    "无效的Token",
	CodeScriptExists:    "脚本已存在",
	CodeScriptNotExists: "脚本不存在",
	CodeInvalidUser:     "用户不对应",
}

func (res ResCode) Msg() string {
	msg, ok := codeMsgMap[res]
	if !ok {
		msg = codeMsgMap[CodeInvalidParam]
	}
	return msg
}
