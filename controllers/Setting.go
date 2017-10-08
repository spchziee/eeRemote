package controllers

import (
	"eePACS/models"

	"github.com/astaxie/beego"
)

type SettingController struct {
	beego.Controller
}

func (this *SettingController) Get() {
	this.Data["Appversion"] = beego.AppConfig.String("appversion")
	models.CurrentUser().SetControllerData(this.Data)
	if models.CurrentUser().IsAdmin() == false {
		this.Redirect("/Home", 301)
		return
	}
	this.TplName = "setting.html"
}
