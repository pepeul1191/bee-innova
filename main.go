package main

import (
	"bee-innova/conf"
	"bee-innova/controllers/common"
	"bee-innova/models/responses"
	_ "bee-innova/routers"
	"encoding/gob"
	"log"
	"os"

	"github.com/beego/beego/v2/server/web"
	beego "github.com/beego/beego/v2/server/web"

	"github.com/joho/godotenv"
)

func registerGobTypes() {
	// Estructuras principales
	gob.Register(responses.LoginResponse{})
	gob.Register(responses.JWTClaims{})

	// Arrays de estructuras
	gob.Register([]responses.RoleWithPermissions{})
	gob.Register([]responses.Permission{})
	gob.Register([]responses.JWTRole{})
	gob.Register([]responses.JWTPermission{})

	// Estructuras individuales
	gob.Register(responses.RoleWithPermissions{})
	gob.Register(responses.Permission{})
	gob.Register(responses.JWTRole{})
	gob.Register(responses.JWTPermission{})

	// Tipos básicos que podrían usarse
	gob.Register(map[string]interface{}{})
	gob.Register([]interface{}{})
}

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
	// Registrar todas las estructuras que puedan ser almacenadas en sesión
	registerGobTypes()
	// prueba .env
	dbPort := os.Getenv("DB_PORT")
	log.Printf("El puerto de la base de datos es: %s", dbPort)
	// Configura el controlador para manejar el error 404
	web.ErrorController(&common.ErrorController{})
	beego.Run()
}
