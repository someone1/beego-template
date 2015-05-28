package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Title"] = "beego.me"
	c.Data["SomeVar"] = "astaxie@gmail.com"
	c.TplNames = "blog/add.tpl"
}
