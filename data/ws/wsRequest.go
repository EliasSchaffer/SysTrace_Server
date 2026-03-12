package ws

type WSRequest struct {
	RequestID string `json:"request_id"`
	Type      string `json:"type"`
	Message   string `json:"message,omitempty"`
	Payload   string `json:"result,omitempty"`
}
