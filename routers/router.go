package routers

import (
	"biggerforum/controllers"

	"github.com/astaxie/beego"
	"github.com/dchest/captcha"
)

func init() {
	beego.Router("/", &controllers.HomeController{})
	beego.Router("/user", &controllers.UserController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/register", &controllers.RegisterController{})
	beego.Router("/question", &controllers.QuestionController{})
	beego.Router("/question/edi", &controllers.QuestionController{}, "get:Edit")
	beego.Router("/files", &controllers.FileController{})

	beego.Handler("/captcha/*", captcha.Server(captcha.StdWidth, captcha.StdHeight))
}
