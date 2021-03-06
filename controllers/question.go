package controllers

import (
	"biggerforum/models"
	"fmt"
)

type QuestionController struct {
	BaseController
}

func (this *QuestionController) URLMapping() {
	this.Mapping("Edit", this.Edit)
}

func (this *QuestionController) Get() {
	res, err := query()
	if err == nil {
		for _, v := range res {
			fmt.Println(v)
		}
	}
}

func (this *QuestionController) Post() {
	title := this.Input().Get("title")
	content := this.Input().Get("content")
	fmt.Println("title :", title)
	fmt.Println("content: ", content)
}

func (this *QuestionController) Edit() {
	this.TplNames = "edit.html"
}

func (this *QuestionController) AddQuestion() {

}

func AddAnswer() {
	question := &models.Question{Id: 1, Title: "我是甩甩"}
	u := &models.User{Id: 1, UserName: "shuaishuai"}
	err := models.AddAnswer(question, u, "傻逼傻逼")
	fmt.Println(err)
}

func query() ([]*models.Question, error) {
	res, err := models.QueryQuestion()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	for _, v := range res {
		loadQuestion(v, models.Relations["Question"])
	}
	return res, nil
}

func loadQuestion(question *models.Question, m []string) {
	for _, v := range m {
		if err := question.LoadRelation(v); err != nil {
			fmt.Println(err)
			continue
		}
	}
}
