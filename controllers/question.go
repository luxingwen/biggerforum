package controllers

import (
	"biggerforum/models"
	"fmt"
)

type QuestionController struct {
	BaseController
}

func (this *QuestionController) Get() {
	//AddQ()
	AddAnswer()
}

func (this *QuestionController) Post() {

}

func AddQ() {
	title := "我是甩甩"
	content := "这个问题不太难"
	u := &models.User{Id: 1, UserName: "shuaishuai"}
	err := models.AddQuestion(title, content, u)
	fmt.Println(err)
}

func AddAnswer() {
	question := &models.Question{Id: 1, Title: "我是甩甩"}
	u := &models.User{Id: 1, UserName: "shuaishuai"}
	err := models.AddAnswer(question, u, "傻逼傻逼")
	fmt.Println(err)
}
