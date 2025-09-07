package services

import (
	"bee-innova/models/responses"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
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

	// Crear request body
	requestBody := map[string]interface{}{
		"username":  username,
		"password":  password,
		"system_id": systemID,
	}

	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		return nil, fmt.Errorf("error marshaling JSON: %v", err)
	}

	// Crear la petición HTTP
	fmt.Println("0 ++++++++++++++++++++++++++++++++")
	fmt.Println(urlAccess + "/api/v1/users/sign-in/by-username")
	req, err := http.NewRequest("POST", urlAccess+"/api/v1/users/sign-in/by-username", bytes.NewBuffer(jsonBody))
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

	// Leer la respuesta completa
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %v", err)
	}

	fmt.Println("1 +++++++++++++++++++++++++++++++++++++++++")
	fmt.Println("Response Body:", string(body))
	fmt.Println("2 +++++++++++++++++++++++++++++++++++++++++")

	// Verificar status code
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("login failed with status: %d - %s", resp.StatusCode, string(body))
	}

	// Decodificar la respuesta
	var apiResponse struct {
		Success bool                    `json:"success"`
		Message string                  `json:"message,omitempty"`
		Data    responses.LoginResponse `json:"data,omitempty"`
		Error   string                  `json:"error,omitempty"`
	}

	if err := json.Unmarshal(body, &apiResponse); err != nil {
		return nil, fmt.Errorf("error decoding JSON response: %v - Response: %s", err, string(body))
	}

	// Verificar si la API retornó éxito
	if !apiResponse.Success {
		return nil, fmt.Errorf("authentication failed: %s", apiResponse.Error)
	}

	// Retornar la respuesta
	return &apiResponse.Data, nil
}
