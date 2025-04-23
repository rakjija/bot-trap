package user

type SignupRequest struct {
	Email    string `json:"email" binding:"required,not_blank,min=1" example:"bob@example.com"`
	Password string `json:"password" binding:"required,not_blank,min=8" example:"pa55w0rd"`
	Username string `json:"username" binding:"required,not_blank" example:"bob"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email" example:"bob@example.com"`
	Password string `json:"password" binding:"required" example:"pa55w0rd"`
}

type MeResponse struct {
	UserID uint `json:"user_id" example:"1"`
}

type SignupResponse struct {
	UserID  uint   `json:"user_id" example:"1"`
	Message string `json:"message" example:"user created successfully"`
}

type LoginResponse struct {
	UserID      uint   `json:"user_id" example:"1"`
	AccessToken string `json:"access_token" example:"eyJhbGciOi..."`
}

type ErrorResponse struct {
	Error string `json:"error" example:"something went wrong"`
}
