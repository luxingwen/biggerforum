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
	ParentId   int64
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

func (this *Topic) Add() error {
	o := orm.NewOrm()
	this.Created = time.Now()
	this.Updated = time.Now()
	_, err := o.Insert(this)
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

func QueryParentTopic() ([]*Topic, error) {
	o := orm.NewOrm()
	var topics []*Topic
	_, err := o.QueryTable("topic").Filter("IsParent", true).All(&topics)
	if err != nil {
		return nil, err
	}
	return topics, nil
}

func GetTopics(t []string) []*Topic {
	o := orm.NewOrm()
	var topics []*Topic
	for _, v := range t {
		t := &Topic{Title: v}
		err := o.Read(t)
		if err != nil {
			continue
		}
		topics = append(topics, t)
	}
	return topics
}

func GetTopic(title string) (*Topic, error) {
	o := orm.NewOrm()
	topic := &Topic{Title: title}
	err := o.Read(topic, "title")
	if err != nil {
		return nil, err
	}
	return topic, nil
}
