package controllers

import (
	"github.com/astaxie/beego"
)

type HelpController struct {
	beego.Controller
}

func (this *HelpController) Get() {
	this.TplName = "help.html"
}
