package controllers

import (
	"biggerforum/models"
	"fmt"
	"strings"
)

type QuestionEdiController struct {
	BaseController
}

func (this *QuestionEdiController) Get() {
	this.TplNames = "edit.html"
	u := this.GetSession("User")
	if u == nil {
		this.Redirect("/", 302)
		return
	}
}

func (this *QuestionEdiController) Post() {
	title := this.Input().Get("title")
	content := this.Input().Get("content")
	topics := this.Input().Get("topics")
	tpoics := strings.Fields(topics)
	tps := models.GetTopics(tpoics)
	u := this.GetSession("User")
	if v, ok := u.(*models.User); ok {
		err := models.AddQuestion(title, content, v, tps)
		if err != nil {
			fmt.Println(err)

		} else {
			this.Redirect("/", 301)
		}

	}

}
