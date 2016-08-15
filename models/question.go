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

func AddQuestion(title, content string, user *User, topics []*Topic) error {
	question := &Question{Title: title, Content: content, User: user, Topics: topics}
	question.Created = time.Now()
	o := orm.NewOrm()
	_, err := o.Insert(question)
	return err
}

func (this *Question) Update() error {
	o := orm.NewOrm()
	this.Updated = time.Now()
	_, err := o.Update(this)
	return err
}

func (this *Question) LoadRelation(table string) error {
	o := orm.NewOrm()
	_, err := o.LoadRelated(this, table)
	return err
}

func QueryQuestion() ([]*Question, error) {
	var questions []*Question
	o := orm.NewOrm()
	_, err := o.QueryTable("question").Limit(100).All(&questions)
	return questions, err
}
