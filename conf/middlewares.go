package conf

import (
	"crypto/rand"
	"encoding/base64"
	"net/http"
	"strings"

	"github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context"
)

const (
	CsrfSessionKey = "_csrf_token"
	CsrfFormKey    = "_csrf"
)

// GenerateToken genera un token CSRF aleatorio
func GenerateToken() string {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return ""
	}
	return base64.URLEncoding.EncodeToString(b)
}

// CSRFMiddleware protege contra ataques CSRF usando tokens de sesión.
func CSRFMiddleware() {
	web.InsertFilter("*", web.BeforeRouter, func(ctx *context.Context) {
		method := ctx.Input.Method()

		// Obtener o crear el token CSRF en la sesión
		session := ctx.Input.CruSession
		sessToken := session.Get(ctx.Request.Context(), CsrfSessionKey)
		if sessToken == nil {
			token := GenerateToken()
			session.Set(ctx.Request.Context(), CsrfSessionKey, token)
			sessToken = token
		}

		// Validar el token solo en métodos sensibles
		if method == http.MethodPost || method == http.MethodPut || method == http.MethodDelete {

			var token string
			contentType := ctx.Input.Header("Content-Type")

			if strings.HasPrefix(contentType, "application/json") {
				token = ctx.Input.Header("X-CSRF-Token")
			} else {
				token = ctx.Input.Query(CsrfFormKey)
			}

			if token == "" || token != sessToken {
				ctx.ResponseWriter.WriteHeader(403)
				ctx.WriteString("Token CSRF inválido o ausente.")
				return
			}
		}
	})
}
