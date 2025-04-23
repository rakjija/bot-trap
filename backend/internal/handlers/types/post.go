package types

type PostCreateRequest struct {
	Title   string `json:"title" binding:"required,min=1" example:"title"`
	Content string `json:"content" binding:"required,min=1" example:"content"`
}

type PostResponse struct {
	ID        uint   `json:"id" example:"1"`
	Title     string `json:"title" example:"title"`
	Content   string `json:"content" example:"content"`
	UserID    uint   `json:"user_id" example:"1"`
	Username  string `json:"username,omitempty" example:"bob (OPTIONAL)"`
	CreatedAt string `json:"created_at" example:"2024-04-23T15:04:05Z"`
}

type UnauthorizedResponse struct {
	Error string `json:"error" example:"unauthorized"`
}

type InternalServerErrorResponse struct {
	Error string `json:"error" example:"something went wrong"`
}
