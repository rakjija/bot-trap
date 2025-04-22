package post

type PostCreateRequest struct {
	Title   string `json:"title" binding:"required,min=1"`
	Content string `json:"content" binding:"required,min=1"`
}

type PostResponse struct {
	ID        uint   `json:"id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	UserID    uint   `json:"user_id"`
	CreatedAt string `json:"created_at"`
}
