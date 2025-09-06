package common

import (
	"github.com/beego/beego/v2/server/web"
)

type LoginController struct {
	web.Controller
}

// ShowLogin muestra el formulario de login.
// @router /login [get]
func (c *LoginController) ShowLogin() {
	c.TplName = "common/login/sign-in.tpl"
}

// Login procesa el formulario de login.
// @router /login [post]
func (c *LoginController) Login() {
	c.TplName = "common/login/sign-in.tpl"
}

// Logout cierra la sesión del usuario.
// Esta es una buena práctica para incluir, aunque no esté en tu ruta de 'common'
// Puedes añadir una ruta en tu 'router.go' como: web.Router("/logout", &common.LoginController{}, "get:Logout")
func (c *LoginController) Logout() {
	// 1. Borrar todos los datos de la sesión
	c.DelSession("user_id")

	// 2. Redirigir al usuario a la página de login
	c.Redirect("/login", 302)
}
