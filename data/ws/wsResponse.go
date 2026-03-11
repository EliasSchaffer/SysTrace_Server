package ws

type WSResponse struct {
	Type    string `json:"type"`
	Message string `json:"message,omitempty"`
	Payload string `json:"result,omitempty"`
}
