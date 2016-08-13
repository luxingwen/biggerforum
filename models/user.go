package models

import (
	"github.com/astaxie/beego/orm"
)

type User struct {
	Id            int64
	UserName      string `orm:"unique"`
	PassWd        string
	Email         string `orm:"unique"`
	VerifyEmail   int
	Phone         string
	Created       int64
	Updated       string
	Lever         int64
	FirstIp       string
	LastIp        string
	FansCount     int
	FlowCount     int
	QuestionCount int
	ArticleCount  int
	AnserCount    int
	GithubUrl     string
	BlogUrl       string
	Question      []*Question      `orm:"null;reverse(many)"`
	TopicFoucs    []*TopicFoucs    `orm:"null;reverse(many)"`
	QuestionFocus []*QuestionFocus `orm:"null;reverse(many)"`
	Answer        []*Answer        `orm:"null;reverse(many)"`
	Fans          []*Fans          `orm:"null;reverse(many)"`
	Flow          []*Flow          `orm:"null;reverse(many)"`
}

type Fans struct {
	Id     int64
	User   *User `orm:"null;rel(fk)"`
	Create int64
}

type Flow struct {
	Id     int64
	User   *User `orm:"null;rel(fk)"`
	Create int64
}

func NewUser(id int64, username, email string) *User {
	return &User{Id: id, UserName: username, Email: email}
}

func Register(username, password, email string) error {
	user := &User{UserName: username, PassWd: password, Email: email}
	o := orm.NewOrm()
	_, err := o.Insert(user)
	return err
}

func Login(username, password string) (*User, error) {
	user := &User{PassWd: password}
	o := orm.NewOrm()
	cond := orm.NewCondition()
	cond1 := cond.Or("user_name", username).Or("email", username)
	cond2 := cond.AndCond(cond1).And("pass_wd", password)
	qs := o.QueryTable("user")
	err := qs.SetCond(cond2).One(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (this *User) SetQuestions() error {
	o := orm.NewOrm()
	num, err := o.QueryTable("question").Filter("user", this.Id).RelatedSel().All(&this.Question)
	if err != nil {
		return err
	}
	this.QuestionCount = int(num)
	return nil
}

func (this *User) SetQuestionFocus() error {
	o := orm.NewOrm()
	_, err := o.QueryTable("question_focus").Filter("user", this.Id).RelatedSel().All(&this.QuestionFocus)
	return err
}

func (this *User) SetAnswer() error {
	o := orm.NewOrm()
	num, err := o.QueryTable("answer").Filter("user", this.Id).RelatedSel().All(&this.Answer)
	if err != nil {
		return err
	}
	this.AnserCount = int(num)
	return nil
}

func (this *User) SetTopicFoucs() error {
	o := orm.NewOrm()
	_, err := o.QueryTable("topic_focus").Filter("user", this.Id).RelatedSel().All(&this.TopicFoucs)
	return err
}

func (this *User) SetFans() error {
	o := orm.NewOrm()
	num, err := o.QueryTable("fans").Filter("user", this.Id).RelatedSel().All(&this.Fans)
	if err != nil {
		return err
	}
	this.FansCount = int(num)
	return nil
}

func (this *User) SetFlow() error {
	o := orm.NewOrm()
	num, err := o.QueryTable("flow").Filter("user", this.Id).RelatedSel().All(&this.Flow)
	if err != nil {
		return err
	}
	this.FlowCount = int(num)
	return nil
}
