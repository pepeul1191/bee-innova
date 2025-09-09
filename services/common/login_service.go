package services

import (
	"bee-innova/models/responses"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

type AuthService struct{}

func NewAuthService() *AuthService {
	return &AuthService{}
}

func (s *AuthService) LoginByUsername(username, password string) (*responses.LoginAPIResponse, error) {
	// Obtener variables de entorno
	urlAccess := os.Getenv("URL_ACCESS_SERVICE")
	if urlAccess == "" {
		return nil, fmt.Errorf("servicio de autenticación no disponible (configuración faltante)")
	}

	xAuthAccess := os.Getenv("X_AUTH_ACCESS_SERVICE")
	if xAuthAccess == "" {
		return nil, fmt.Errorf("servicio de autenticación no disponible (credenciales faltantes)")
	}

	systemIDStr := os.Getenv("SYSTEM_ID")
	if systemIDStr == "" {
		return nil, fmt.Errorf("servicio de autenticación no disponible (sistema no configurado)")
	}

	systemID, err := strconv.Atoi(systemIDStr)
	if err != nil {
		return nil, fmt.Errorf("configuración del sistema inválida: %v", err)
	}

	// Validar credenciales antes de enviar
	if username == "" || password == "" {
		return nil, fmt.Errorf("usuario y contraseña son requeridos")
	}

	// Crear request body
	requestBody := map[string]interface{}{
		"username":  username,
		"password":  password,
		"system_id": systemID,
	}

	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		return nil, fmt.Errorf("error interno al procesar la solicitud")
	}

	// Crear la petición HTTP
	req, err := http.NewRequest("POST", urlAccess+"/api/v1/users/sign-in/by-username", bytes.NewBuffer(jsonBody))
	if err != nil {
		log.Printf("Error creando petición HTTP: %v", err)
		return nil, fmt.Errorf("error de conexión con el servicio de autenticación")
	}

	// Agregar headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("X-Auth-Trigger", xAuthAccess)

	// Realizar petición
	client := &http.Client{
		Timeout: 30 * time.Second, // Agregar timeout para evitar bloqueos
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Error en petición HTTP: %v", err)
		return nil, fmt.Errorf("No se pudo conectar con el servicio de autenticación")
	}
	defer resp.Body.Close()

	// Leer la respuesta completa
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error leyendo respuesta: %v", err)
		return nil, fmt.Errorf("Error al procesar la respuesta del servicio")
	}

	// Log para debugging (opcional en producción)
	log.Printf("Respuesta del servicio de autenticación: %s", string(body))

	// Decodificar la respuesta
	var apiResponse responses.LoginAPIResponse
	if err := json.Unmarshal(body, &apiResponse); err != nil {
		log.Printf("Error decodificando JSON: %v - Respuesta: %s", err, string(body))
		return nil, fmt.Errorf("respuesta inválida del servicio de autenticación")
	}

	return &apiResponse, nil
}
