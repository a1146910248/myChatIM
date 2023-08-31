package controllers

type LongPollingController struct {
	baseController
}

func (c *LongPollingController) Join() {
	// safe check
	uname := c.GetString("uname")
	if len(uname) == 0 {
		c.Redirect("/", 302)
		return
	}

	//join room
	//Join(uname, nil)

	c.TplName = "longpolling.html"
	c.Data["IsLongPolling"] = true
	c.Data["UserName"] = uname
}
