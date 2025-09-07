package routers

import (
	"bee-innova/controllers/common"

	"github.com/beego/beego/v2/server/web"
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

	web.AddNamespace(commonNs)
}
