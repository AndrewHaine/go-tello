package web

import (
	"time"

	"github.com/andrewhaine/go-tello/pkg/tello"
)

type EventType string

const (
	EventTypeLog EventType = "log.created"
	EventTypeCommand EventType = "command.requested"
	EventTypeConnection EventType = "connection.updated"
	EventTypeTelemetry EventType = "telemetry.updated"
)

type Event struct {
	event EventType
	payload map[string]any
	timestamp time.Time
}

type CommandEvent struct {
	Event
	payload map[string]struct{command string}
}

func EventFromTelemetry(telemetry tello.Telemetry) Event {
	return Event{
		event: EventTypeTelemetry,
		payload: map[string]any{
			"battery": telemetry.Bat,
			"pitch": telemetry.Pitch,
			"roll": telemetry.Roll,
			"yaw": telemetry.Yaw,
			"temp_high": telemetry.Temp.High,
			"temp_low": telemetry.Temp.Low,
			"height": telemetry.Height,
		},
		timestamp: time.Now(),
	}
}

func EventFromTelloMsg(msg tello.Message) Event {
	return Event{
		event: EventTypeLog,
		payload: map[string]any{
			"message": msg.Message,
			"time": msg.Time,
		},
		timestamp: time.Now(),
	}
}

func EventFromConnectionStatus(connected bool) Event {
	connectionStatus := "connected"

	if !connected {
		connectionStatus = "disconnected"
	}

	return Event{
		event: EventTypeConnection,
		payload: map[string]any{
			"connection_status": connectionStatus,
		},
		timestamp: time.Now(),
	}
}
