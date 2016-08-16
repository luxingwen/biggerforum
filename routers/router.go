package routers

import (
	"biggerforum/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.HomeController{})
	beego.Router("/user", &controllers.UserController{})
	beego.Router("/user/login", &controllers.UserController{}, "get:Login;post:LoginIn")
	beego.Router("/user/register", &controllers.UserController{}, "get:Register;post:RegisterIn")
	beego.Router("/question", &controllers.QuestionController{})
	beego.Router("/question/edi", &controllers.QuestionController{}, "get:Edit")
	beego.Router("/files", &controllers.FileController{})
}
