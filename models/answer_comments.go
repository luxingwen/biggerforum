package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type AnswerComments struct {
	Answer   *Answer `orm:"null;rel(fk)"`
	FromUser *User   `orm:"null;rel(one)`
	ToUser   *User   `orm:"null;rel(one)"`
	Content  string
	Created  time.Time
}
