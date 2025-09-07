package main

import (
	"bee-innova/conf"
	"bee-innova/controllers/common"
	_ "bee-innova/routers"
	"log"
	"os"

	"github.com/beego/beego/v2/server/web"
	beego "github.com/beego/beego/v2/server/web"

	"github.com/joho/godotenv"
)

func main() {
	// .env
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error al cargar el archivo .env")
	}

	// csrf
	web.BConfig.WebConfig.Session.SessionOn = true

	conf.CSRFMiddleware()

	// helpers
	web.AddFuncMap("styles", conf.GenerateStylesHTML)
	web.AddFuncMap("scripts", conf.GenerateScriptsHTML)
	web.AddFuncMap("GetEnv", conf.GetEnv)
	// prueba .env
	dbPort := os.Getenv("DB_PORT")
	log.Printf("El puerto de la base de datos es: %s", dbPort)
	// Configura el controlador para manejar el error 404
	web.ErrorController(&common.ErrorController{})
	beego.Run()
}
