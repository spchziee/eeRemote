package controllers

import (
	"eePACS/models"

	"github.com/astaxie/beego"
)

type AdminController struct {
	beego.Controller
}

func (this *AdminController) Get() {
	this.Data["Appversion"] = beego.AppConfig.String("appversion")
	models.CurrentUser().SetControllerData(this.Data)
	if models.CurrentUser().IsLogin() == false {
		this.Redirect("/", 301)
		return
	}
	this.TplName = "admin.html"
}
