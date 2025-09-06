package common

import (
	"github.com/beego/beego/v2/server/web"
)

type WebController struct {
	web.Controller
}

// ShowLogin muestra el formulario de login.
// @router /login [get]
func (c *WebController) ShowHome() {
	c.TplName = "common/web/index.tpl"
}
