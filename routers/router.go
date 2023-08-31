package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"myChatIM/controllers"
)

func init() {
	// 欢迎界面
	beego.Router("/", &controllers.AppController{})
	// 提交登录表单
	beego.Router("/join", &controllers.AppController{}, "post:Join")

	// Long Polling界面
	beego.Router("/lp", &controllers.LongPollingController{}, "get:Join")
}
