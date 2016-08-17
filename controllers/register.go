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

}

func (this *RegisterController) RegisterIn() {
	username := this.Input().Get("username")
	password := this.Input().Get("password")
	email := this.Input().Get("email")
	user, err := models.Register(username, password, email)
	if err != nil {
		fmt.Println(err)
		this.Redirect("/user/register", 302)
	}
	this.Redirect("/", 302)
	this.SetSession("User", user)
}
