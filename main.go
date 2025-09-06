package main

import (
	_ "bee-innova/routers"
	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	beego.Run()
}

