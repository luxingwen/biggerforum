package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type Question struct {
	Id           int64
	Title        string
	Content      string
	User         *User `orm:"null;rel(fk)"`
	AnswerCount  int
	Answer       []*Answer `orm:"null;reverse(many)"`
	ViewCount    int
	FocusCount   int
	Topics       []*Topic `orm:"null;rel(m2m)"`
	AgreeId      int
	AgainstCount int64
	BestAnswerId int64
	Created      time.Time `orm:"null;index"`
	Updated      time.Time `orm:"null;index"`
}

func AddQuestion(title, content string, user *User) error {
	question := &Question{Title: title, Content: content, User: user}
	question.Created = time.Now()
	o := orm.NewOrm()
	_, err := o.Insert(question)
	return err
}

func (this *Question) Update() error {
	o := orm.NewOrm()
	_, err := o.Update(this)
	return err
}

func (this *Question) SetTopics() error {
	o := orm.NewOrm()
	_, err := o.QueryTable("question_topics").Filter("question_id", this.Id).All(&this.Topics)
	return err
}

func (this *Question) SetAnswer() error {
	o := orm.NewOrm()
	num, err := o.QueryTable("answer").Filter("answer_id", this.Id).All(&this.Topics)
	if err != nil {
		return err
	}
	this.AnswerCount = int(num)
	return nil
}

func (this *Question) SetUser() error {
	o := orm.NewOrm()
	err := o.QueryTable("user").Filter("question_id", this.Id).Limit(1).One(&this.User)
	return err
}
