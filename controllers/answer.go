package controllers

import (
	"biggerforum/models"
	"log"
)

type AnswerController struct {
	BaseController
}

func (this *AnswerController) Get() {

}

func loadAnswer(this *models.Answer, m []string) {
	for _, v := range m {
		if err := this.LoadRelation(v); err != nil {
			log.Println(err)
			continue
		}
	}
}
