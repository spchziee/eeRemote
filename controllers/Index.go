package controllers

import (
	"github.com/astaxie/beego"
)

type IndexController struct {
	beego.Controller
}

func (this *IndexController) Get() {
	redirect := this.GetString("redirect")
	if len(redirect) > 0 {
		this.Redirect("/"+redirect, 301)
		return
	}

	this.Data["Email"] = "Info@mednova.cn"
	this.TplName = "index.html"
}
