// controllers/common/login_controller.go
package common

import (
	"bee-innova/conf"
	"bee-innova/controllers"
	"bee-innova/helpers/common"
)

type LoginController struct {
	controllers.BaseController
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

// @router /sign-in [post]
func (c *LoginController) Login() {
	username := c.GetString("username")
	password := c.GetString("password")

	if username == "" || password == "" {
		c.SetFlash("warning", "Campos no pueden estar vacíos")
		c.Redirect("/login", 302)
		return
	}

	// Llamar al servicio de autenticación
	response, err := controllers.AuthService.LoginByUsername(username, password)
	if err != nil {
		// Manejar error del servicio
		c.SetFlash("danger", "Error en el servicio de autenticación: "+err.Error())
		c.Redirect("/sign-in", 302)
		return
	}

	c.SetSession("auth_token", response.Token)
	c.SetSession("user_id", response.UserID)
	c.SetFlash("success", "¡Login exitoso!")
	c.SetSession("username", username)
	c.Redirect("/sign-in", 302)
}

func (c *LoginController) Logout() {
	// 1. Borrar todos los datos de la sesión
	c.DelSession("user_id")

	// 2. Redirigir al usuario a la página de login
	c.Redirect("/login", 302)
}
