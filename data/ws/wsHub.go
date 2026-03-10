package ws

// WSDirectMessage kapselt eine Nachricht an genau einen Client.
type WSDirectMessage struct {
	ClientID string
	Message  []byte
	Result   chan bool
}

type WSHub struct {
	Clients     map[*WSClient]bool
	ClientsByID map[string]*WSClient
	Register    chan *WSClient
	Unregister  chan *WSClient
	Broadcast   chan []byte
	DirectSend  chan WSDirectMessage
}

func (h *WSHub) Run() {
	for {
		select {
		case client := <-h.Register:
			h.Clients[client] = true
			if client.ClientID != "" {
				h.ClientsByID[client.ClientID] = client
			}
		case client := <-h.Unregister:
			if _, ok := h.Clients[client]; ok {
				delete(h.Clients, client)
				if client.ClientID != "" {
					if mapped, found := h.ClientsByID[client.ClientID]; found && mapped == client {
						delete(h.ClientsByID, client.ClientID)
					}
				}
				close(client.Send)
			}
		case message := <-h.Broadcast:
			for client := range h.Clients {
				select {
				case client.Send <- message:
				default:
					close(client.Send)
					delete(h.Clients, client)
					if client.ClientID != "" {
						if mapped, found := h.ClientsByID[client.ClientID]; found && mapped == client {
							delete(h.ClientsByID, client.ClientID)
						}
					}
				}
			}
		case direct := <-h.DirectSend:
			ok := false
			if client, found := h.ClientsByID[direct.ClientID]; found {
				select {
				case client.Send <- direct.Message:
					ok = true
				default:
					close(client.Send)
					delete(h.Clients, client)
					delete(h.ClientsByID, direct.ClientID)
				}
			}
			if direct.Result != nil {
				direct.Result <- ok
			}
		}
	}
}
