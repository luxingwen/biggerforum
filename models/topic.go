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
	Updated    time.Time
	Questions  []*Question `orm:"null;reverse(many)"` //设置一对多的反向关系
}

func AddTopic(title, desrc, pic string) error {
	topit := &Topic{Title: title, Descr: desrc, Pic: pic, Created: time.Now()}
	o := orm.NewOrm()
	_, err := o.Insert(topit)
	return err
}

func (this *Topic) Update() error {
	o := orm.NewOrm()
	this.Updated = time.Now()
	_, err := o.Update(this)
	return err
}

func (this *Topic) LoadRelation(table string) error {
	o := orm.NewOrm()
	_, err := o.LoadRelated(this, table)
	return err
}

func QueryTopic() ([]*Topic, error) {
	o := orm.NewOrm()
	var topics []*Topic
	_, err := o.QueryTable("topic").All(&topics)
	return topics, err
}
