package controllers

import (
	"biggerforum/models"
	"fmt"
)

type UserController struct {
	BaseController
}

func (this *UserController) Get() {
	//	username := "shuaishuai"
	//	password := "123456"
	//email := "935232474@qq.com"
	//Login(username, password)
	QueryQuestions()
}

func (this *UserController) Post() {

}

func (this *UserController) Put() {

}

func Register(username, password, email string) {
	err := models.Register(username, password, email)
	fmt.Println(err)
}
func Login(username, password string) {
	u, err := models.Login(username, password)
	fmt.Println(u, err)
}

func QueryQuestions() {
	u := models.NewUser(1, "shuaishuai", "935232474@qq.com")
	err := u.SetQuestions()
	fmt.Println(err, *u)
}
