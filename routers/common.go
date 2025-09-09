// routers/common.go
package routers

import (
	"bee-innova/controllers/common"

	"github.com/beego/beego/v2/server/web"
)

func CommonRoutes() {
	web.Router("/sign-in", &common.LoginController{}, "get:ShowSignIn;post:Login")
	web.Router("/", &common.WebController{}, "get:ShowHome")
	web.Router("/api/v1/session", &common.LoginController{}, "get:GetSessionData")
	web.Router("/about", &common.WebController{}, "get:ShowAbout")
	web.Router("/contact", &common.WebController{}, "get:ShowContact")
}
