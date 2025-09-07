package responses

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	SystemID int    `json:"system_id"`
}

type LoginResponse struct {
	Token   string `json:"token,omitempty"`
	UserID  int    `json:"user_id,omitempty"`
	Message string `json:"message,omitempty"`
	Success bool   `json:"success"`
}
