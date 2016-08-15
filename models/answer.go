package models

import (
	"errors"
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

func (this *Answer) LoadRelation(table string) error {
	o := orm.NewOrm()
	_, err := o.LoadRelated(this, table)
	return err
}

func (this *Answer) Delete(u *User) error {
	o := orm.NewOrm()
	if this.User.Id != u.Id {
		return errors.New("no permission!")
	}
	_, err := o.Delete(this)
	return err
}
