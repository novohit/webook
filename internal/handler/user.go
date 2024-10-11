package handler

import (
	"net/http"
	"regexp"

	regexp2 "github.com/dlclark/regexp2"

	"github.com/gin-gonic/gin"
)

const (
	emailRegexPattern = "^\\w+([-+.]\\w+)*@\\w+([-.]\\w+)*\\.\\w+([-.]\\w+)*$"
	// 和上面比起来，用 ` 看起来就比较清爽
	passwordRegexPattern = `^(?=.*[A-Za-z])(?=.*\d)(?=.*[$@$!%*#?&])[A-Za-z\d$@$!%*#?&]{8,}$`
)

type UserHandler struct {
	passwordExp *regexp2.Regexp
}

func NewUserHandler() *UserHandler {
	//optimize 预编译正则表达式
	re := regexp2.MustCompile(passwordRegexPattern, 0)
	return &UserHandler{passwordExp: re}
}

func (u *UserHandler) SignIn(ctx *gin.Context) {
	//
}

func (u *UserHandler) SignUp(ctx *gin.Context) {
	type SignUpReq struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	var req SignUpReq
	// Bind 会返回400 ShouldBindJSON 需要自己设置
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if matched, _ := regexp.Match(emailRegexPattern, []byte(req.Email)); !matched {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "email format error"})
		return
	}

	// go regexp 不支持部分复杂的语法 需要第三方库
	// err 可能是因为超时 不要忽略
	matched, err := u.passwordExp.MatchString(req.Password)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if !matched {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "password format error"})
		return
	}

	ctx.JSON(200, gin.H{"message": "success"})
}

func (u *UserHandler) Profile(ctx *gin.Context) {
	//
}

func (u *UserHandler) Edit(ctx *gin.Context) {

}