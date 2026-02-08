package web

import (
	"fmt"
	"log"
	"time"

	"github.com/gorilla/websocket"
)

type Browser struct {
	Conn *websocket.Conn
	Hub *Hub
	Queue chan Event
}

const (
	pongWait = 60 * time.Second
	pingPeriod = (pongWait * 9) / 10
	maxMessageSize = 512
)

func (b *Browser) ReceiveCommands() {
	defer func() {
		b.Hub.Deregister <- b
		b.Conn.Close()
	}()

	b.Conn.SetReadLimit(maxMessageSize)
	b.Conn.SetReadDeadline(time.Now().Add(pongWait))
	b.Conn.SetPongHandler(func(string) error { b.Conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })

	for {
		var cmdEvent CommandEvent
		err := b.Conn.ReadJSON(&cmdEvent)

		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("Error reading from browser: %v", err)
			}
			fmt.Println(err.Error())
			break
		}

		if cmdEvent.Event == EventTypeCommand {
			b.Hub.Commands <- []byte(cmdEvent.Payload.Command)
		}
	}
}

func (b *Browser) SendQueuedMessages() {
	ticker := time.NewTicker(pingPeriod)

	defer func () {
		ticker.Stop()
		b.Conn.Close()
	}()

	for {
		select {
		case msg, ok := <- b.Queue:
			if !ok {
				b.Conn.WriteMessage(websocket.CloseMessage, nil)
				return
			}

			b.Conn.WriteJSON(msg)
		case <-ticker.C:
			b.Conn.WriteMessage(websocket.PingMessage, nil)
		}
	}
}
