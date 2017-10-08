package controllers

import (
	"eePACS/models"

	"github.com/astaxie/beego"
)

type PatientListController struct {
	beego.Controller
}

func (this *PatientListController) Get() {
	this.Data["Appversion"] = beego.AppConfig.String("appversion")
	models.CurrentUser().SetControllerData(this.Data)
	if models.CurrentUser().IsLogin() == false {
		this.Redirect("/", 301)
		return
	}
	chktypes := [...]string{"DX", "CT", "CR", "MR"}
	repdoctors := [...]string{"DX", "CT", "CR", "MR"}
	auddoctors := [...]string{"DX", "CT", "CR", "MR"}

	this.Data["PatName"] = this.GetString("patname")
	this.Data["PatID"] = this.GetString("patID")
	this.Data["ChkTypes"] = chktypes
	this.Data["ChkType"] = this.GetString("chktype")
	this.Data["RepState"] = this.GetString("repstate")
	this.Data["RepDoctors"] = repdoctors
	this.Data["RepDoctor"] = this.GetString("repdoctor")
	this.Data["AudDoctors"] = auddoctors
	this.Data["AudDoctor"] = this.GetString("auddoctor")

	this.Data["StayCode"] = this.GetString("staycode")
	this.Data["ClinicCode"] = this.GetString("cliniccode")

	this.Data["StartChkDate"] = this.GetString("startchkdate")
	this.Data["EndChkDate"] = this.GetString("endchkdate")
	this.Data["StartRepDate"] = this.GetString("startrepdate")
	this.Data["EndRepDate"] = this.GetString("endrepdate")
	this.Data["StartAudDate"] = this.GetString("startauddate")
	this.Data["EndAudDate"] = this.GetString("endauddate")
	this.Data["KeyWord"] = this.GetString("keyword")
	if this.GetString("arrange") != "1" {
		this.Data["Arrange1"] = "selected"
		this.Data["Arrange2"] = ""
	} else {
		this.Data["Arrange1"] = ""
		this.Data["Arrange2"] = "selected"
	}

	this.TplName = "patlist.html"
}
