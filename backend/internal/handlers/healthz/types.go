package healthz

type HealthResponse struct {
	Status string `json:"status" example:"ok"`
}
