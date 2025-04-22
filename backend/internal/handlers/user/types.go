package user

type SignupRequest struct {
	Email    string `json:"email" binding:"required,not_blank,min=1"`
	Password string `json:"password" binding:"required,not_blank,min=8"`
	Username string `json:"username" binding:"required,not_blank"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}
