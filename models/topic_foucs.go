package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type TopicFoucs struct {
	Id      int64
	Topic   *Topic `orm:"null;rel(one)"` //
	User    *User  `orm:"null;rel(fk)"`
	Created time.Time
}

func (this *TopicFoucs) LoadRelation(table string) error {
	o := orm.NewOrm()
	_, err := o.LoadRelated(this, table)
	return err
}

func (this *TopicFoucs) Add() error {
	o := orm.NewOrm()
	this.Created = time.Now()
	_, err := o.Insert(this)
	return err
}
