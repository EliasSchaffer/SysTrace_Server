package ws

type WSResponse struct {
	Type      string `json:"type"`
	RequestID string `json:"request_id"`
	Status    int    `json:"status"`
	Message   string `json:"message,omitempty"`
}
