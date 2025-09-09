// filters/redirect-if-authenticated.go
package filters

import (
	"github.com/beego/beego/v2/server/web/context"
)

func RedirectIfAuthenticated(ctx *context.Context) {
	// Verificar si existe sesión de usuario o roles
	if ctx.Input.Session("user") != nil || ctx.Input.Session("roles") != nil {
		// Redirigir a la página principal
		ctx.Redirect(302, "/")
	}
}

func RedirectIfAuthenticatedOnGet(ctx *context.Context) {
	// Solo procesar para método GET
	if ctx.Request.Method != "GET" {
		return
	}

	// Verificar si existe sesión de usuario o roles
	if ctx.Input.Session("user") != nil || ctx.Input.Session("roles") != nil {
		// Redirigir a la página principal
		ctx.Redirect(302, "/")
	}
}
