package types

type LogPayload struct {
	Timestamp string `json:"timestamp"`
	Duration  string `json:"duration,omitempty"`
	Level     string `json:"level"`
	Service   string `json:"service"`
	Method    string `json:"method"`
	Path      string `json:"path"`
	Status    int    `json:"status"`
	IP        string `json:"ip"`
	UserAgent string `json:"user_agent"`
	Message   string `json:"message"`
	User      string `json:"user,omitempty"`
}
