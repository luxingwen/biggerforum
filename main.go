package main

import (
	"biggerforum/models"
	_ "biggerforum/routers"
	"fmt"

	"github.com/astaxie/beego/orm"

	"github.com/astaxie/beego"
)

func init() {
	models.RegisterDb()
}

func main() {
	orm.Debug = true
	err := orm.RunSyncdb("default", false, true)
	if err != nil {
		fmt.Println(err)
	}
	beego.Run()
}
