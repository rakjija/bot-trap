package db

import "gorm.io/gorm"

type LogEntry struct {
	gorm.Model
	IP      string `json:"ip" binding:"required"`
	Path    string `json:"path" binding:"required"`
	Message string `json:"message" binding:"required"`
}
