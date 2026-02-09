package web

import (
	"time"

	"github.com/andrewhaine/go-tello/pkg/tello"
	"github.com/pion/webrtc/v3"
)

type EventType string

const (
	EventTypeLog EventType = "log.created"
	EventTypeCommand EventType = "command.requested"
	EventTypeConnection EventType = "connection.updated"
	EventTypeTelemetry EventType = "telemetry.updated"
	EventTypeVideoPeerConnOffer EventType = "video.offer"
	EventTypeVideoPeerConnAnswer EventType = "video.answer"
	EventTypeVideoPeerConnIceCandidate EventType = "video.ice-candidate"
	EventTypeScreenshotAdded EventType = "screenshot.added"
)

type Event struct {
	Event EventType `json:"event"`
	Payload map[string]any `json:"payload"`
	Timestamp time.Time `json:"timestamp"`
}

type CommandEvent struct {
	Event EventType `json:"event"`
	Payload struct{
		Command string `json:"command"`
	} `json:"payload"`
	Timestamp time.Time `json:"timestamp"`
}

type VideoPeerConnOfferEvent struct {
	Event EventType `json:"event"`
	Payload webrtc.SessionDescription `json:"payload"`
}

type VideoPeerConnAnswerEvent struct {
	Event EventType `json:"event"`
	Payload map[string]any `json:"payload"`
}

type VideoPeerConnIceCandidateEvent struct {
	Event EventType `json:"event"`
	Payload map[string]any `json:"payload"`
}

type ScreenshotAddedEvent struct {
	Event EventType `json:"event"`
	Payload map[string]any `json:"payload"`
}

func EventFromTelemetry(telemetry tello.Telemetry) Event {
	return Event{
		Event: EventTypeTelemetry,
		Payload: map[string]any{
			"battery": telemetry.Bat,
			"pitch": telemetry.Pitch,
			"roll": telemetry.Roll,
			"yaw": telemetry.Yaw,
			"temp_high": telemetry.Temp.High,
			"temp_low": telemetry.Temp.Low,
			"height": telemetry.Height,
		},
		Timestamp: time.Now(),
	}
}

func EventFromTelloMsg(msg tello.Message) Event {
	return Event{
		Event: EventTypeLog,
		Payload: map[string]any{
			"message": msg.Message,
			"time": msg.Time,
		},
		Timestamp: time.Now(),
	}
}

func EventFromConnectionStatus(connected bool) Event {
	connectionStatus := "connected"

	if !connected {
		connectionStatus = "disconnected"
	}

	return Event{
		Event: EventTypeConnection,
		Payload: map[string]any{
			"connection_status": connectionStatus,
		},
		Timestamp: time.Now(),
	}
}
