package main

import (
	_ "myChatIM/routers"
	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	beego.Run()
}

