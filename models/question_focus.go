package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type QuestionFocus struct {
	Id       int64
	Question *Question `orm:"null;rel(one)"`
	User     *User     `orm:"null;rel(fk)"`
	Created  time.Time
}

func (this *QuestionFocus) Add() error {
	this.Created = time.Now()
	o := orm.NewOrm()
	_, err := o.Insert(this)
	return err
}
