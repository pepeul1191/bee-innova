package responses

import (
	"github.com/golang-jwt/jwt/v5"
)

// LoginResponse - Estructura principal de respuesta del login
type LoginResponse struct {
	User  UserLoginResponse     `json:"user"`
	Roles []RoleWithPermissions `json:"roles"`
}

// UserLoginResponse - Información del usuario en la respuesta de login
type UserLoginResponse struct {
	ID       uint64 `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Token    string `json:"token"`
}

// RoleWithPermissions - Rol con sus permisos
type RoleWithPermissions struct {
	ID          uint64       `json:"id"`
	Name        string       `json:"name"`
	Permissions []Permission `json:"permissions"`
}

// Permission - Estructura de permiso
type Permission struct {
	ID   uint64 `json:"id"`
	Name string `json:"name"`
}

// Claims personalizados para el JWT (si necesitas decodificar el token)
type JWTClaims struct {
	UserID   uint64    `json:"user_id"`
	Username string    `json:"username"`
	Email    string    `json:"email"`
	Roles    []JWTRole `json:"roles"`
	jwt.RegisteredClaims
}

// JWTRole - Rol dentro del JWT
type JWTRole struct {
	ID          uint64          `json:"id"`
	Name        string          `json:"name"`
	Permissions []JWTPermission `json:"permissions"`
}

// JWTPermission - Permiso dentro del JWT
type JWTPermission struct {
	ID   uint64 `json:"id"`
	Name string `json:"name"`
}

type LoginAPIResponse struct {
	Success bool          `json:"success"`
	Message string        `json:"message,omitempty"`
	Data    LoginResponse `json:"data,omitempty"`
	Error   string        `json:"error,omitempty"`
}
