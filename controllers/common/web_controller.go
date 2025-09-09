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
	// Verificar si el usuario está logueado
	if user := c.GetSession("user"); user != nil {
		// Usuario logueado - mostrar dashboard o vista personalizada
		c.Data["PageTitle"] = "Dashboard - Bienvenido"
		c.Data["Styles"] = common.GetDashboardStylesHelper()
		c.Data["Scripts"] = common.GetDashboardScriptsHelper()
		c.Data["Navlink"] = "dashboard"
		c.Data["User"] = user // Pasar datos del usuario a la vista

		c.Layout = "layouts/application.tpl"
		c.TplName = "common/dashboard/index.tpl"
	} else {
		// Usuario no logueado - vista normal
		c.Data["PageTitle"] = "Bienvenido a InnovaULima"
		c.Data["Styles"] = common.GetIndexStylesHelper()
		c.Data["Scripts"] = common.GetIndexScriptsHelper()
		c.Data["Navlink"] = "index"

		c.Layout = "layouts/web.tpl"
		c.TplName = "common/web/index.tpl"
	}
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
