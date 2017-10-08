package controllers

import (
	"fmt"

	"github.com/astaxie/beego"
)

type UserController struct {
	beego.Controller
}

type User struct {
	UserName   string `form:"username"`
	FamilyName string `form:"familyname"`
	Pass       string `form:"pass"`
}

func (this *UserController) Login() {
	user := &User{}
	this.ParseForm(user)
	fmt.Println(user)
	this.SetSession("login", int(1))
	this.Redirect("/Home", 301)
	return
}

func (this *UserController) Logout() {
	user := &User{}
	this.ParseForm(user)
	fmt.Println(user)
	this.DelSession("login")
	this.Redirect("/", 301)
	return
}

func (this *UserController) Register() {
	user := &User{}
	this.ParseForm(user)
	fmt.Println(user)
	this.SetSession("login", int(1))
	this.Redirect("/Home", 301)
	return
}

func (this *UserController) UsernameExist() {
	fmt.Println(this.GetString("v"))
	if this.GetString("v") == "ipxipx" {
		this.Ctx.ResponseWriter.Write([]byte("true"))
	} else {
		this.Ctx.ResponseWriter.Write([]byte("false"))
	}

}

/*
beego中引入jwt权限认证,有效的实行SSO(单点登录) http://blog.csdn.net/hzwy23/article/details/53314306
*/
