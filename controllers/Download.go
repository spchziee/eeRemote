package controllers

import (
	"github.com/astaxie/beego"
)

type DownloadController struct {
	beego.Controller
}

var mapDownloadFile map[string]string

func init() {
	mapDownloadFile = make(map[string]string)
	mapDownloadFile["DCMViewer.zip"] = "./static/无法浏览图像.zip"
}

func (this *DownloadController) Get() {
	filename := this.Ctx.Input.Param(":splat")
	if v, ok := mapDownloadFile[filename]; ok {
		this.Ctx.Output.Download(v)
	} else {
		this.Redirect("/404", 301)
	}
}
