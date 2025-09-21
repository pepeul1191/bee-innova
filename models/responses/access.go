package responses

import (
	"github.com/golang-jwt/jwt/v5"
)

// AccessResponse - Estructura principal de respuesta del login
type AccessAPIResponse struct {
	Success bool               `json:"success"`
	Message string             `json:"message"`
	Error   string             `json:"error"`
	Data    AccessDataResponse `json:"data"`
}

type AccessDataResponse struct {
	User  UserAccessResponse    `json:"user"`
	Roles []RoleWithPermissions `json:"roles"`
	Token string                `json:"token"`
}

// FilesResponse - Estructura principal de respuesta del login
type FilesAPIResponse struct {
	Success bool              `json:"success"`
	Message string            `json:"message"`
	Error   string            `json:"error"`
	Data    FilesDataResponse `json:"data"`
}

type FilesDataResponse struct {
	Token string `json:"token"`
}

// LoginResponse - Estructura principal de respuesta del login
type LoginResponse struct {
	Success bool                  `json:"success"`
	Message string                `json:"message"`
	User    UserAccessResponse    `json:"user"`
	Roles   []RoleWithPermissions `json:"roles"`
	Tokens  Tokens                `json:"tokens"`
}

// Permission - Estructura de permiso
type Tokens struct {
	Access string `json:"access"`
	File   string `json:"files"`
}

// UserAccessResponse - Información del usuario en la respuesta de login
type UserAccessResponse struct {
	ID       uint64 `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
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
