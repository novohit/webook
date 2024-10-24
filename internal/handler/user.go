package handler

import (
	"net/http"
	"regexp"
	"webook/internal/domain"
	"webook/internal/service"

	regexp2 "github.com/dlclark/regexp2"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

const (
	emailRegexPattern = "^\\w+([-+.]\\w+)*@\\w+([-.]\\w+)*\\.\\w+([-.]\\w+)*$"
	// 和上面比起来，用 ` 看起来就比较清爽
	passwordRegexPattern = `^(?=.*[A-Za-z])(?=.*\d)(?=.*[$@$!%*#?&])[A-Za-z\d$@$!%*#?&]{8,}$`
)

type UserHandler struct {
	svc         *service.UserService
	passwordExp *regexp2.Regexp
}

func NewUserHandler(svc *service.UserService) *UserHandler {
	//optimize 预编译正则表达式
	re := regexp2.MustCompile(passwordRegexPattern, 0)
	return &UserHandler{
		svc:         svc,
		passwordExp: re,
	}
}

func (u *UserHandler) SignIn(ctx *gin.Context) {
	type SignInReq struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	var req SignInReq
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := u.svc.SignIn(ctx, domain.User{
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}
	session := sessions.Default(ctx)
	session.Set("user_id", user.Id)
	session.Options(sessions.Options{
		MaxAge: 30 * 60,
		// 线上环境需要配置
		//Secure:   true,
		//HttpOnly: true,
	})
	session.Save()
	ctx.JSON(200, gin.H{"message": "success", "data": user})
}

func (u *UserHandler) SignInJWT(ctx *gin.Context) {
	type SignInReq struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	var req SignInReq
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := u.svc.SignInJWT(ctx, domain.User{
		Email:    req.Email,
		Password: req.Password,
	}, ctx.Request.UserAgent())
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"message": "success", "data": token})
}

func (u *UserHandler) SignOut(ctx *gin.Context) {
	session := sessions.Default(ctx)
	session.Options(sessions.Options{MaxAge: -1})
	// MaxAge 设置成-1 redis里会自动清空对应的session 不需要手动删除
	//session.Delete("user_id")
	session.Save()
	ctx.JSON(200, gin.H{"message": "success"})
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

	err = u.svc.SignUp(ctx, domain.User{Email: req.Email, Password: req.Password})
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"message": "success"})
}

func (u *UserHandler) Profile(ctx *gin.Context) {
	//
	value, _ := ctx.Get("identify")
	ctx.JSON(http.StatusOK, gin.H{"message": "success", "data": value})
}

func (u *UserHandler) Edit(ctx *gin.Context) {

}
