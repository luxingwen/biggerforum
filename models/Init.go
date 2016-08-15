package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

var (
	Relations map[string][]string
)

func init() {
	Relations = map[string][]string{
		"User":           {"Question", "TopicFoucs", "QuestionFocus", "Answer", "Fans", "Flow"},
		"Topic":          {"Questions"},
		"TopicFoucs":     {"Topic", "User"},
		"Question":       {"User", "Answer", "Topics"},
		"QuestionFoucs":  {"Question", "User"},
		"Answer":         {"Question", "User", "AnswerComments"},
		"AnswerComments": {"Answer", "FromUser", "ToUser"},
	}
}
func RegisterDb() {
	orm.RegisterModel(new(User), new(Answer), new(TopicFoucs), new(Fans), new(Question), new(Topic), new(QuestionFocus), new(Flow), new(AnswerComments))
	orm.RegisterDriver("mysql", orm.DR_MySQL)
	orm.RegisterDataBase("default", "mysql", "root:root@tcp(127.0.0.1:3306)/bbs?charset=utf8")
}
