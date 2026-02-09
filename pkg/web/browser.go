package web

import (
	"encoding/json"
	"log"
	"time"

	"github.com/gorilla/websocket"
	"github.com/pion/webrtc/v3"
)

type Browser struct {
	Conn *websocket.Conn
	Hub *Hub
	VideoPeerConn *webrtc.PeerConnection
	Queue chan Event
}

const (
	pongWait = 60 * time.Second
	pingPeriod = (pongWait * 9) / 10
	maxMessageSize = 12000
)

func (b *Browser) ReceiveMessages() {
	defer func() {
		b.Hub.Deregister <- b
		b.Conn.Close()
	}()

	b.Conn.SetReadLimit(maxMessageSize)
	b.Conn.SetReadDeadline(time.Now().Add(pongWait))
	b.Conn.SetPongHandler(func(string) error { 
		b.Conn.SetReadDeadline(time.Now().Add(pongWait))
		return nil 
	})
	
	peerConn, _ := webrtc.NewPeerConnection(webrtc.Configuration{})
	b.VideoPeerConn = peerConn
	b.VideoPeerConn.AddTrack(b.Hub.VideoServer.VideoTrack)

	b.VideoPeerConn.OnICECandidate(func(candidate *webrtc.ICECandidate) {
		if candidate != nil {
			b.Conn.WriteJSON(map[string]any{
				"event": "video.ice-candidate",
				"payload": candidate.ToJSON(),
			})
		}
	})

	offer, _ := b.VideoPeerConn.CreateOffer(nil)
	b.VideoPeerConn.SetLocalDescription(offer)

	offerMsg := VideoPeerConnOfferEvent{
		Event: EventTypeVideoPeerConnOffer,
		Payload: offer,
	}
	b.Conn.WriteJSON(offerMsg)

	for {
		var event Event
		err := b.Conn.ReadJSON(&event)
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("Error reading from browser: %v", err)
			}
			break
		}

		if event.Event == EventTypeCommand {
			if cmd, ok := event.Payload["command"].(string); ok {
				b.Hub.Commands <- []byte(cmd)
			}
		}

		if event.Event == EventTypeVideoPeerConnAnswer {
			var answer webrtc.SessionDescription
			data, _ := json.Marshal(event.Payload)
			json.Unmarshal(data, &answer)
			b.VideoPeerConn.SetRemoteDescription(answer)
		}

		if event.Event == EventTypeVideoPeerConnIceCandidate {
			var candidate webrtc.ICECandidateInit
			data, _ := json.Marshal(event.Payload)
			json.Unmarshal(data, &candidate)
			b.VideoPeerConn.AddICECandidate(candidate)
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
