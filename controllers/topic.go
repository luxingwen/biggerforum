package controllers

import (
	"biggerforum/models"
	"fmt"
)

type TopicController struct {
	BaseController
}

func (this *TopicController) Get() {

}

func loadTopic(topic *models.Topic, m []string) {
	for _, v := range m {
		if err := topic.LoadRelation(v); err != nil {
			fmt.Println(err)
			continue
		}
	}
}
