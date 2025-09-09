// controllers/common/login_controller.go
package common

import (
	"bee-innova/conf"
	"bee-innova/controllers"
	"bee-innova/helpers/common"
	"net/http"
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
		c.SetFlash("danger", err.Error())
		c.Redirect("/sign-in", 302)
		return
	}

	if response.Success {
		c.SetSession("user", response.Data.User)
		c.SetSession("roles", response.Data.Roles)
		c.SetFlash("success", "¡Login exitoso!")
	} else {
		c.SetFlash("danger", response.Message)
	}
	c.Redirect("/sign-in", 302)
}

func (c *LoginController) Logout() {
	// 1. Borrar todos los datos de la sesión
	c.DelSession("user_id")

	// 2. Redirigir al usuario a la página de login
	c.Redirect("/login", 302)
}

func (c *LoginController) GetSessionData() {
	// Obtener datos de la sesión
	userData := c.GetSession("user")
	rolesData := c.GetSession("roles")

	// Verificar si el usuario está autenticado
	if userData == nil {
		c.Ctx.Output.SetStatus(http.StatusUnauthorized)
		c.Data["json"] = map[string]interface{}{
			"success": false,
			"message": "Usuario no autenticado",
			"data":    nil,
		}
		c.ServeJSON()
		return
	}

	// Preparar respuesta
	response := map[string]interface{}{
		"success": true,
		"message": "Datos de sesión obtenidos correctamente",
		"data": map[string]interface{}{
			"user":  userData,
			"roles": rolesData,
		},
	}

	c.Data["json"] = response
	c.ServeJSON()
}
