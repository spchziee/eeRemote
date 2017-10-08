package main

import (
	"eePACS/fileTrans"
	_ "eePACS/routers"
	"eePACS/spcUtils"

	"github.com/astaxie/beego"
)

func main() {

	go fileTrans.Start("10.10.10.160:4567")

	beego.AddFuncMap("SetSelected", spcUtils.SetSelected)

	beego.BConfig.WebConfig.Session.SessionOn = true
	//beego.BConfig.WebConfig.EnableXSRF = true
	beego.BConfig.WebConfig.XSRFKey = "mednova_eepacs"
	beego.BConfig.WebConfig.XSRFExpire = 3600

	beego.Run()
}
