package controllers

import (
	"biggerforum/models"
	"fmt"
)

type UserController struct {
	BaseController
}

func (this *UserController) Get() {

}

func (this *UserController) Post() {

}

func (this *UserController) Put() {

}

func loadUser(user *models.User, m []string) {
	for _, v := range m {
		if err := user.LoadRelation(v); err != nil {
			fmt.Println(err)
			continue
		}
	}
}
