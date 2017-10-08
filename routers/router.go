package routers

import (
	"eePACS/controllers"
	"html/template"
	"net/http"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.IndexController{})

	//用户相关方法
	beego.Router("/U/Login", &controllers.UserController{}, "post:Login")
	beego.Router("/U/Logout", &controllers.UserController{}, "get:Logout")
	beego.Router("/U/Register", &controllers.UserController{}, "post:Register")
	beego.Router("/U/UsernameExist", &controllers.UserController{}, "get:UsernameExist")

	//页面
	beego.Router("/Home", &controllers.HomeController{})
	beego.Router("/Help", &controllers.HelpController{})
	beego.Router("/Download/*", &controllers.DownloadController{})
	beego.Router("/PatList", &controllers.PatientListController{})
	beego.Router("/Admin", &controllers.AdminController{})
	beego.Router("/DcmView", &controllers.DcmViewController{})
	beego.Router("/PrintSetting", &controllers.PrintSettingController{})
	beego.Router("/Setting", &controllers.SettingController{})
	beego.Router("/Statistics", &controllers.StatisticsController{})

	beego.ErrorHandler("404", func(rw http.ResponseWriter, r *http.Request) {
		t, _ := template.New("404.html").ParseFiles("./views/404.html")
		t.Execute(rw, nil)
	})

}
