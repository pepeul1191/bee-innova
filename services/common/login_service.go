package services

import (
	"bee-innova/models/responses"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

type AuthService struct{}

func NewAuthService() *AuthService {
	return &AuthService{}
}

func (s *AuthService) LoginByUsername(username, password string) (*responses.LoginResponse, error) {
	// Obtener variables de entorno
	urlAccess := os.Getenv("URL_ACCESS_SERVICE")
	if urlAccess == "" {
		return nil, fmt.Errorf("URL_ACCESS_SERVICE environment variable not set")
	}

	xAuthAccess := os.Getenv("X_AUTH_ACCESS_SERVICE")
	if xAuthAccess == "" {
		return nil, fmt.Errorf("X_AUTH_ACCESS_SERVICE environment variable not set")
	}

	systemIDStr := os.Getenv("SYSTEM_ID")
	if systemIDStr == "" {
		return nil, fmt.Errorf("SYSTEM_ID environment variable not set")
	}

	systemID, err := strconv.Atoi(systemIDStr)
	if err != nil {
		return nil, fmt.Errorf("invalid SYSTEM_ID: %v", err)
	}

	// Crear request body con username en lugar de email
	requestBody := responses.LoginRequest{
		Username: username,
		Password: password,
		SystemID: systemID,
	}

	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		return nil, fmt.Errorf("error marshaling JSON: %v", err)
	}

	// Crear la petición HTTP
	req, err := http.NewRequest("POST", urlAccess+"/api/v1/sign-in/by-username", bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	// Agregar headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("X-Auth-Trigger", xAuthAccess)
	// Realizar petición
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making request: %v", err)
	}
	defer resp.Body.Close()
	// Verificar status code
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("login failed with status: %s", resp.Status)
	}
	// Leer respuesta
	var loginResp responses.LoginResponse
	if err := json.NewDecoder(resp.Body).Decode(&loginResp); err != nil {
		return nil, fmt.Errorf("error decoding response: %v", err)
	}

	loginResp.Success = true
	return &loginResp, nil
}
