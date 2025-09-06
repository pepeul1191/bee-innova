package routers

import (
	"bee-innova/controllers"
	"bee-innova/controllers/common"

	"github.com/beego/beego/v2/server/web"
	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	// cargar otros routers
	CommonRoutes()

	// montar otros routers
	commonNs := web.NewNamespace("/",
		web.NSInclude(
			&common.LoginController{},
			// ... otros controladores
		),
	)

	beego.Router("/demo", &controllers.MainController{})
	web.AddNamespace(commonNs)
}
