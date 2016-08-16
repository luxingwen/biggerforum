package controllers

import (
	"biggerforum/models"
	"fmt"
)

type UserController struct {
	BaseController
}

func (this *UserController) URLMapping() {
	this.Mapping("Login", this.Login)
	this.Mapping("LoginIn", this.LoginIn)
	this.Mapping("Register", this.Register)
	this.Mapping("RegisterIn", this.RegisterIn)
}

func (this *UserController) Get() {
}

func (this *UserController) Post() {

}

func (this *UserController) Put() {

}

func (this *UserController) Login() {
	this.TplNames = "login.html"
}

func (this *UserController) LoginIn() {
	username := this.Input().Get("username")
	password := this.Input().Get("password")
	user, err := login(username, password)
	if err != nil {
		fmt.Println("err: ", err)
		this.Redirect("/user/login", 302)
	}
	this.SetSession("User", user)
	this.Redirect("/", 301)
}
func (this *UserController) Register() {
	this.TplNames = "register.html"
}
func (this *UserController) RegisterIn() {
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

func login(username, password string) (*models.User, error) {
	u, err := models.Login(username, password)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func loadUser(user *models.User, m []string) {
	for _, v := range m {
		if err := user.LoadRelation(v); err != nil {
			fmt.Println(err)
			continue
		}
	}
}
