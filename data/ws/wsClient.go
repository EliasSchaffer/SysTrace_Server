package ws

import (
	"fmt"

	"github.com/gorilla/websocket"
)

type WSClient struct {
	WSHub           *WSHub
	ClientID        string
	Conn            *websocket.Conn
	Send            chan []byte
	OnClientMessage func(clientID string, message []byte)
}

func (w *WSClient) ReadPump() {
	defer func() {
		w.WSHub.Unregister <- w
		w.Conn.Close()
	}()

	for {
		_, message, err := w.Conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				fmt.Printf("Error: %v", err)
			}
			break
		}
		fmt.Printf("recv: %s", message)
		if w.OnClientMessage != nil {
			w.OnClientMessage(w.ClientID, message)
		}
	}
}

func (c *WSClient) WritePump() {
	defer func() {
		c.Conn.Close()
	}()

	for {
		message, ok := <-c.Send
		if !ok {
			c.Conn.WriteMessage(websocket.CloseMessage, []byte{})
			return
		}

		err := c.Conn.WriteMessage(websocket.TextMessage, message)
		if err != nil {
			return
		}
	}
}
