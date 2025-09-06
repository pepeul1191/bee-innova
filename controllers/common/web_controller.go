// controllers/common/web_controller.go
package common

import (
	"github.com/beego/beego/v2/server/web"
)

type WebController struct {
	web.Controller
}

// @router / [get]
func (c *WebController) ShowHome() {
	c.Layout = "layouts/blank.tpl"
	c.Data["PageTitle"] = "Bienvenido a InnovaULima"
	c.Data["Styles"] = []string{"main", "login", "forms"}
	c.Data["Scripts"] = []string{"jquery", "app"}
	c.TplName = "common/web/index.tpl"
}
