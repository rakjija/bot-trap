package types

type LogPayload struct {
	Level   string `json:"level"`
	Message string `json:"message"`
	User    string `json:"user,omitempty"`
}
