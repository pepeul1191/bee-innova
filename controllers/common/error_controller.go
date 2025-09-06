package common

import (
	"bee-innova/helpers/common"
	"net/http"

	"github.com/beego/beego/v2/server/web"
)

// ErrorController gestiona los errores de la aplicación.
type ErrorController struct {
	web.Controller
}

// Error404 gestiona los errores HTTP 404.
func (c *ErrorController) Error404() {
	// Comprueba si la petición viene de un cliente que espera JSON (ej. una API)
	if c.Ctx.Input.Header("Accept") == "application/json" {
		c.Ctx.Output.SetStatus(http.StatusNotFound)
		c.Data["json"] = map[string]string{
			"error":   "Not Found",
			"message": "The requested resource was not found.",
		}
		c.ServeJSON()
		return
	}

	// Para cualquier otro caso (navegadores web), renderiza la vista HTML
	c.Data["Styles"] = common.GetError404StylesHelper()
	c.Layout = "layouts/blank.tpl"
	c.TplName = "common/error/404.tpl"
}
