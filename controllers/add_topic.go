package controllers

import (
	"biggerforum/models"
	"fmt"
)

type AddTopicController struct {
	BaseController
}

func (this *AddTopicController) Get() {
	this.TplNames = "addTopic.html"
	tps, err := models.QueryParentTopic()
	if err != nil {
		fmt.Println("query error:", err)
	}
	this.Data["Topics"] = tps
}

func (this *AddTopicController) Post() {
	topic := new(models.Topic)
	title := this.Input().Get("title")
	parent := this.Input().Get("isParent") == "on"
	topic.IsParent = true
	if !parent {
		fatherTopic := this.Input().Get("fatherTopic")
		parentTopic, err := models.GetTopic(fatherTopic)
		if err != nil {
			fmt.Println("GetTopic error:", err)
			return
		}
		topic.ParentId = parentTopic.Id
		topic.IsParent = false
	}
	topic.Title = title
	err := topic.Add()
	if err != nil {
		fmt.Println(err)
	}
	this.Redirect("/topic/add", 301)
}
