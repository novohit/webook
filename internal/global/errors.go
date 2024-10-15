package global

import "errors"

var (
	ErrUserOrPassword = errors.New("账号或密码输入错误")
	ErrUserNotFound   = errors.New("用户不存在")
)
