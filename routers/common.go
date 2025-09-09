// routers/common.go
package routers

import (
	"bee-innova/controllers/common"
	"bee-innova/filters"

	"github.com/beego/beego/v2/server/web"
)

func CommonRoutes() {
	// Aplicar filtro para redirigir usuarios autenticados
	web.InsertFilter("/sign-in", web.BeforeRouter, filters.RedirectIfAuthenticatedOnGet)
	web.InsertFilter("/about", web.BeforeRouter, filters.RedirectIfAuthenticated)
	web.InsertFilter("/contact", web.BeforeRouter, filters.RedirectIfAuthenticated)
	web.InsertFilter("/api/v1/session", web.BeforeRouter, filters.RequireAuth)
	web.InsertFilter("/session", web.BeforeRouter, filters.RequireAuth)

	web.Router("/", &common.WebController{}, "get:ShowHome")
	web.Router("/sign-in", &common.LoginController{}, "get:ShowSignIn;post:Login")
	web.Router("/sign-out", &common.LoginController{}, "get:Logout")
	web.Router("/api/v1/session", &common.LoginController{}, "get:GetSessionData")
	web.Router("/session", &common.LoginController{}, "get:GetSessionData")
	web.Router("/about", &common.WebController{}, "get:ShowAbout")
	web.Router("/contact", &common.WebController{}, "get:ShowContact")
	web.Router("/error/access/401", &common.ErrorController{}, "get:Error401")
}
