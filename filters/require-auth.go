// filters/require-auth.go
package filters

import (
	"strings"

	"github.com/beego/beego/v2/server/web/context"
)

func RequireAuth(ctx *context.Context) {
	// Verificar si NO existe sesión de usuario o roles
	if ctx.Input.Session("user") == nil && ctx.Input.Session("roles") == nil {
		// Determinar el tipo de respuesta basado en Accept header o método
		acceptHeader := ctx.Request.Header.Get("Accept")
		isAPIRequest := strings.Contains(acceptHeader, "application/json") ||
			strings.Contains(ctx.Request.URL.Path, "/api/")

		// Para solicitudes GET que no son API, mostrar página de acceso denegado
		if ctx.Request.Method == "GET" && !isAPIRequest {
			// Redirigir a página de acceso no autorizado
			ctx.Redirect(302, "/error/access/401")
			return
		}

		// Para otros métodos (POST, PUT, DELETE) o solicitudes API, retornar JSON
		ctx.Output.SetStatus(401)
		ctx.Output.JSON(map[string]interface{}{
			"success": false,
			"message": "Acceso no autorizado. Por favor inicie sesión.",
			"error":   "unauthorized",
		}, false, false)
		return
	}
}
