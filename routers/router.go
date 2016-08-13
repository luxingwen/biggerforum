package routers

import (
	"biggerforum/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.BaseController{})
	beego.Router("/user", &controllers.UserController{})
	beego.Router("/question", &controllers.QuestionController{})
}
