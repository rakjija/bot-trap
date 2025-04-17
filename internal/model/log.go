package model

type LogRequest struct {
	IP      string `json:"ip" binding:"required"`
	Path    string `json:"path" binding:"required"`
	Message string `json:"message" binding:"required"`
}
