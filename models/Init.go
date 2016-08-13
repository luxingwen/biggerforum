package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func RegisterDb() {
	orm.RegisterModel(new(User), new(Answer), new(TopicFoucs), new(Fans), new(Question), new(Topic), new(QuestionFocus), new(Flow))
	orm.RegisterDriver("mysql", orm.DR_MySQL)
	orm.RegisterDataBase("default", "mysql", "root:root@tcp(127.0.0.1:3306)/bbs?charset=utf8")
}
