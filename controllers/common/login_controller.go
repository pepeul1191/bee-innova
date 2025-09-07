// controllers/common/login_controller.go
package common

import (
	"bee-innova/conf"
	"bee-innova/helpers/common"
	"fmt"
)

type LoginController struct {
	BaseController
}

// @router /sign-in [get]
func (c *LoginController) ShowSignIn() {
	csrfToken := conf.GetCSRFToken(c.Ctx)

	c.Data["PageTitle"] = "Bienvenido"
	c.Data["Styles"] = common.GetLoginStylesHelper()
	c.Data["Navlink"] = "about"
	c.Data["CsrfToken"] = csrfToken
	c.Layout = "layouts/blank.tpl"
	c.TplName = "common/login/sign-in.tpl"
}

func (c *LoginController) Login() {
	username := c.GetString("username")
	password := c.GetString("password")

	if username == "" || password == "" {
		c.SetFlash("warning", "Campos no pueden estar vacíos")
		c.Redirect("/login", 302)
		return
	}

	// Lógica de autenticación...
	if username != "admin" || password != "password123" {
		fmt.Println("Credenciales incorrectas.")
		c.SetFlash("danger", "Credenciales incorrectas.")
		c.Redirect("/sign-in", 302)
		return
	}

	c.SetFlash("success", "¡Login exitoso!")
	c.SetSession("username", username)
	c.Redirect("/sign-in", 302)
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
