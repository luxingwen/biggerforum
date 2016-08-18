package controllers

import (
	"biggerforum/models"
	"fmt"
	"net/http"

	"github.com/dchest/captcha"
)

type LoginController struct {
	BaseController
}

func (this *LoginController) Get() {
	this.TplNames = "login.html"
	this.Data["CaptchaId"] = captcha.New()
}

func (this *LoginController) Post() {
	email := this.Input().Get("email")
	password := this.Input().Get("password")
	if !verifyCode(this.Ctx.Request) {
		this.Ctx.WriteString("verifycode error")
		return
	}
	user, err := login(email, password)
	if err != nil {
		fmt.Println("err: ", err)
		this.Redirect("/user", 302)
	}
	this.SetSession("User", user)
	this.Redirect("/", 301)
}

func login(username, password string) (*models.User, error) {
	u, err := models.Login(username, password)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func verifyCode(req *http.Request) bool {
	return captcha.VerifyString(req.FormValue("captchaId"), req.FormValue("captchaSolution"))
}
