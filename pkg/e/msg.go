package e

import "errors"

const (
	SUCCESS            = "ok"
	ERROR              = "fail"
	InvalidParams      = "请求参数错误"
	MissingJwt         = "缺少JWT"
	BrokenJwt          = "JWT有误"
	UserNotFound       = "找不到该用户"
	WrongPassword      = "密码错误"
	AuthCheckTokenFail = "解析Token出错"
)

var (
	ErrUserNotFound = errors.New(UserNotFound)
)
