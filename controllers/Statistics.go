package controllers

import (
	"eePACS/models"

	"github.com/astaxie/beego"
)

type StatisticsController struct {
	beego.Controller
}

func (this *StatisticsController) Get() {
	this.Data["Appversion"] = beego.AppConfig.String("appversion")
	models.CurrentUser().SetControllerData(this.Data)
	if models.CurrentUser().IsAdmin() == false {
		this.Redirect("/Home", 301)
		return
	}
	this.TplName = "statistics.html"
}
