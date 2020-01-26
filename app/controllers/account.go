package controllers

import (
	"net/http"

	"github.com/chalvern/apollo/app/service"
	"github.com/chalvern/apollo/configs/initializer"
	"github.com/chalvern/sugar"
	"github.com/gin-gonic/gin"
)

// SigninGet 获取登录页面
func SigninGet(c *gin.Context) {
	c.HTML(http.StatusOK, "account/signin.tpl", gin.H{
		"PageTitle": "登录",
	})
}

// SignupGet 获取注册页面
func SignupGet(c *gin.Context) {
	c.HTML(http.StatusOK, "account/signup.tpl", gin.H{
		"PageTitle": "注册",
	})
}

// SigninPost 注册
func SigninPost(c *gin.Context) {
	pageTitle := "注册"
	form := struct {
		Email     string `form:"email" binding:"required,email,lenlte=50"`
		Password  string `form:"password" binding:"required,lengte=8"`
		Password2 string `form:"password2" binding:"required,gtefield=Password,ltefield=Password"`
		CaptchaID string `form:"captcha_id" binding:"required"`
		Captcha   string `form:"captcha" binding:"required"`
	}{}
	// https://github.com/go-playground/validator/tree/v8.18.2
	if errs := c.ShouldBind(&form); errs != nil {
		sugar.Warnf("SigninPost Bind form Error: %s", errs.Error())
		// errors := errs.(validator.ValidationErrors)
		c.HTML(http.StatusOK, "account/signup.tpl", gin.H{
			"PageTitle": pageTitle,
			FlashError:  "请检查邮箱、密码、验证码内容及格式是否填写正确",
		})
		return
	}

	// 验证码校验
	if !initializer.Captcha.Verify(form.CaptchaID, form.Captcha) {
		c.HTML(http.StatusBadRequest, "account/signup.tpl", gin.H{
			"PageTitle": pageTitle,
			FlashError:  "验证码错误",
		})
		return
	}

	if err := service.UserSignup(form.Email, form.Password); err != nil {
		c.HTML(http.StatusBadRequest, "account/signup.tpl", gin.H{
			"PageTitle": pageTitle,
			FlashError:  "创建用户失败，邮箱已注册",
		})
		return
	}

	htmlOfOk(c, "notify/success.tpl", pageTitle, gin.H{
		"Info":         "注册成功 😆😆😆",
		"Timeout":      5,
		"RedirectURL":  "/signin",
		"RedirectName": "登陆页",
	})

}
