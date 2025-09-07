// controllers/common/web_controller.go
package common

import (
	"bee-innova/helpers/common"

	"github.com/beego/beego/v2/server/web"
)

type WebController struct {
	web.Controller
}

// @router / [get]
func (c *WebController) ShowHome() {
	//dbPort := os.Getenv("DB_PORT")
	//log.Printf("WebController: El puerto de la base de datos es: %s", dbPort)
	c.Data["PageTitle"] = "Bienvenido a InnovaULima"
	c.Data["Styles"] = common.GetIndexStylesHelper()
	c.Data["Scripts"] = common.GetIndexScriptsHelper()
	c.Data["Navlink"] = "index"
	c.Layout = "layouts/web.tpl"
	c.TplName = "common/web/index.tpl"
}

// @router /about [get]
func (c *WebController) ShowAbout() {
	c.Data["PageTitle"] = "Acerca de Nosotros"
	c.Data["Styles"] = common.GetAboutStylesHelper()
	c.Data["Scripts"] = common.GetIndexScriptsHelper()
	c.Data["Navlink"] = "about"
	c.Layout = "layouts/web.tpl"
	c.TplName = "common/web/about.tpl"
}

// @router /contact [get]
func (c *WebController) ShowContact() {
	c.Data["PageTitle"] = "Contáctanos"
	c.Data["Styles"] = common.GetContactStylesHelper()
	c.Data["Scripts"] = common.GetIndexScriptsHelper()
	c.Data["Navlink"] = "contact"
	c.Layout = "layouts/web.tpl"
	c.TplName = "common/web/contact.tpl"
}
