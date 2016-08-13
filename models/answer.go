package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type Answer struct {
	Id             int64
	Question       *Question `orm:"null;rel(fk)"` //设置一对多关系
	Content        string
	Created        time.Time
	User           *User             `orm:"null;rel(fk)"` //设置一对多关系
	AnswerComments []*AnswerComments `orm:"null;reverse(many)"`
}

func AddAnswer(question *Question, user *User, content string) error {
	answer := &Answer{Question: question, User: user, Content: content, Created: time.Now()}
	o := orm.NewOrm()
	_, err := o.Insert(answer)
	return err
}

func (this *Answer) SetUser() error {
	o := orm.NewOrm()
	err := o.QueryTable("user").Filter("answer_id", this.id).Limit(1).One(&this.User)
	return err
}

func (this *Answer) SetAnswerComments() error {
	o := orm.NewOrm()
	_, err := o.QueryTable("answer_comments").Filter("answer_id", this.Id).All(&this.AnswerComments)
	return err
}
