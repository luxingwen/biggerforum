package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type Topic struct {
	Id         int64
	Title      string
	Descr      string
	Pic        string
	FocusCount int
	ParentId   int
	IsParent   bool
	Created    time.Time
	Questions  []*Question `orm:"null;reverse(many)"` //设置一对多的反向关系
}

type TopicFoucs struct {
	Id      int64
	Topic   *Topic `orm:"null;rel(one)"` //
	User    *User  `orm:"null;rel(fk)"`
	Created time.Time
}

func AddTopic(title, desrc string) error {
	topit := &Topic{Title: title, Descr: desrc, Created: time.Now()}
	o := orm.NewOrm()
	_, err := o.Insert(topit)
	return err
}

func (this *Topic) Update() error {
	o := orm.NewOrm()
	_, err := o.Update(this)
	return err
}
