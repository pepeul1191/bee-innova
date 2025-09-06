// controllers/common/web_controller.go
package common

import (
	"bee-innova/helpers/common"
	"log"
	"os"

	"github.com/beego/beego/v2/server/web"
)

type WebController struct {
	web.Controller
}

// @router / [get]
func (c *WebController) ShowHome() {
	dbPort := os.Getenv("DB_PORT")
	log.Printf("WebController: El puerto de la base de datos es: %s", dbPort)
	c.Layout = "layouts/blank.tpl"
	c.Data["PageTitle"] = "Bienvenido a InnovaULima"
	c.Data["Styles"] = common.GetIndexStylesHelper()
	c.Data["Scripts"] = common.GetIndexScriptsHelper()
	c.TplName = "common/web/index.tpl"
}
