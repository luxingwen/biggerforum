package controllers

import (
	"biggerforum/models"
	"fmt"

	"github.com/dchest/captcha"
)

type RegisterController struct {
	BaseController
}

func (this *RegisterController) Get() {
	this.TplNames = "register.html"
	this.Data["CaptchaId"] = captcha.New()
}

func (this *RegisterController) Post() {
	username := this.Input().Get("username")
	password := this.Input().Get("password")
	email := this.Input().Get("email")
	if !verifyCode(this.Ctx.Request) {
		this.Ctx.WriteString("verifycode error")
		return
	}
	user, err := models.Register(username, password, email)
	if err != nil {
		fmt.Println(err)
		this.Redirect("/register", 302)
	}
	this.Redirect("/", 302)
	this.SetSession("User", user)
}
