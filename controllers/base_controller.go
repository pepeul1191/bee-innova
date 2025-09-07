package controllers

import (
	services "bee-innova/services/common"
	"fmt"

	"github.com/beego/beego/v2/server/web"
)

var AuthService = services.NewAuthService()

type BaseController struct {
	web.Controller
}

// SetFlash establece un mensaje flash directamente en la sesión
func (c *BaseController) SetFlash(level, message string) {
	// Crear mapa con los datos del flash
	flashData := map[string]interface{}{
		"level":   level,
		"message": message,
	}

	// Guardar directamente en la sesión
	c.SetSession("flash", flashData)

	fmt.Printf("Flash guardado en sesión: level=%s, message=%s\n", level, message)
}

// GetFlash obtiene el mensaje flash de la sesión
func (c *BaseController) GetFlash() map[string]interface{} {
	flashData := c.GetSession("flash")
	if flashData == nil {
		return nil
	}

	// Convertir a mapa
	if flashMap, ok := flashData.(map[string]interface{}); ok {
		return flashMap
	}

	return nil
}

// HasFlash verifica si hay mensajes flash en la sesión
func (c *BaseController) HasFlash() bool {
	return c.GetSession("flash") != nil
}

// ClearFlash limpia el mensaje flash de la sesión
func (c *BaseController) ClearFlash() {
	c.DelSession("flash")
	fmt.Println("Flash limpiado de sesión")
}

// Render sobreescrito para manejar flash messages automáticamente
func (c *BaseController) Render() error {
	// Procesar flash messages antes de renderizar
	if c.HasFlash() {
		c.Data["Flash"] = c.GetFlash()
		c.ClearFlash() // Limpiar después de asignar
	}

	return c.Controller.Render()
}
