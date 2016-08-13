package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type AnswerComments struct {
	Id       int64
	Answer   *Answer `orm:"null;rel(fk)"`
	FromUser *User   `orm:"null;rel(one)"`
	ToUser   *User   `orm:"null;rel(one)"`
	Content  string
	Created  time.Time
}

func AddAnswerComments(answer *Answer, fromUser, toUser *User) error {
	an := &AnswerComments{Answer: answer, FromUser: fromUser, ToUser: toUser}
	o := orm.NewOrm()
	_, err := o.Insert(an)
	return err
}
