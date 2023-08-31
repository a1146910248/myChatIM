package controllers

import beego "github.com/beego/beego/v2/server/web"

func init() {

}

type baseController struct {
	beego.Controller // 继承beego controller

	//TODO 国际化
}

// AppController handles the welcome screen that allows user to pick a technology and username.
type AppController struct {
	baseController
}

func (c *AppController) Get() {
	c.TplName = "welcome.html"
}

func (c *AppController) Join() {
	// Get from value
	uname := c.GetString("uname")
	tech := c.GetString("tech")

	// check valid
	if len(uname) == 0 {
		c.Redirect("/", 302)
		return
	}
	switch tech {
	case "longpolling":
		c.Redirect("/lp?uname="+uname, 302)
	case "websocket":
		c.Redirect("/wb?uname="+uname, 302)
	default:
		c.Redirect("/", 302)
	}
	return
}
