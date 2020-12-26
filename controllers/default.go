package controllers

import (
	"demo/models"
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	models.UpdatePage()
	page := models.GetPage()
	c.Data["Website"] = page.Website
	c.Data["Email"] = page.Email
	c.TplName = "index.tpl"
}
