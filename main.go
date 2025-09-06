package main

import (
	"bee-innova/conf"
	_ "bee-innova/routers"

	"github.com/beego/beego/v2/server/web"
	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	// helpers
	web.AddFuncMap("styles", conf.GenerateStylesHTML)
	web.AddFuncMap("scripts", conf.GenerateScriptsHTML)

	beego.Run()
}
