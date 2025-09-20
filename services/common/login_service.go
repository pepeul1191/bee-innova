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

// AuthService handles authentication-related operations.
type AuthService struct{}

// NewAuthService creates and returns a new AuthService instance.
func NewAuthService() *AuthService {
	return &AuthService{}
}

// LoginByUsername handles user login by username and password.
func (s *AuthService) LoginByUsername(username, password string) (*responses.LoginResponse, error) {

	// Validate input
	if username == "" || password == "" {
		return nil, fmt.Errorf("username and password are required")
	}

	// Call the private helper function to perform the HTTP request
	accessAPIResponse, err := s.makeAuthRequest(username, password)
	if err != nil {
		return nil, err
	}

	// Map the response to the desired format
	var loginResponse responses.LoginResponse
	loginResponse.Success = accessAPIResponse.Success
	loginResponse.Message = accessAPIResponse.Message
	loginResponse.User = accessAPIResponse.Data.User
	loginResponse.Roles = accessAPIResponse.Data.Roles
	loginResponse.Tokens = []responses.TokenJWT{
		{Name: "access", JWT: accessAPIResponse.Data.Token},
	}

	return &loginResponse, nil
}

// makeAuthRequest is a private helper function that encapsulates the HTTP request logic.
func (s *AuthService) makeAuthRequest(username, password string) (*responses.AccessAPIResponse, error) {
	// 1. Get environment variables
	urlAccess := os.Getenv("URL_ACCESS_SERVICE")
	xAuthHeader := os.Getenv("X_AUTH_ACCESS_SERVICE")
	systemIDStr := os.Getenv("SYSTEM_ID")

	if urlAccess == "" || xAuthHeader == "" || systemIDStr == "" {
		return nil, fmt.Errorf("authentication service is not available (missing configuration)")
	}

	systemID, err := strconv.Atoi(systemIDStr)
	if err != nil {
		return nil, fmt.Errorf("invalid system configuration: %v", err)
	}

	requestBody := map[string]interface{}{
		"username":  username,
		"password":  password,
		"system_id": systemID,
	}

	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		log.Printf("Internal error marshalling JSON: %v", err)
		return nil, fmt.Errorf("internal error processing the request")
	}

	// Create the HTTP request
	req, err := http.NewRequest("POST", urlAccess+"/api/v1/users/sign-in/by-username", bytes.NewBuffer(jsonBody))
	if err != nil {
		log.Printf("Error creating HTTP request: %v", err)
		return nil, fmt.Errorf("connection error with the authentication service")
	}

	// Add headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("X-Auth-Trigger", xAuthHeader)

	// Perform the request
	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Error during HTTP request: %v", err)
		return nil, fmt.Errorf("could not connect to the authentication service")
	}
	defer resp.Body.Close()

	// Read the full response body
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error reading response: %v", err)
		return nil, fmt.Errorf("error processing the service response")
	}

	log.Printf("Authentication service response: %s", string(respBody))

	// Decode the response
	var accessAPIResponse responses.AccessAPIResponse
	if err := json.Unmarshal(respBody, &accessAPIResponse); err != nil {
		log.Printf("Error decoding JSON: %v - Response: %s", err, string(respBody))
		return nil, fmt.Errorf("invalid response from the authentication service")
	}

	return &accessAPIResponse, nil
}
