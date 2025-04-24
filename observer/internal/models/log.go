package models

type LogEntry struct {
	Timestamp string  `json:"timestamp"`
	Duration  float64 `json:"duration,omitempty"`
	Level     string  `json:"level"`
	Service   string  `json:"service"`
	Method    string  `json:"method"`
	Path      string  `json:"path"`
	Status    int     `json:"status"`
	IP        string  `json:"ip"`
	UserAgent string  `json:"user_agent"`
	Message   string  `json:"message"`
	User      string  `json:"user,omitempty"`
}
