package tui

import (
	"fmt"

	"github.com/andrewhaine/go-tello/pkg/tello"
)

type Vitals struct {
  bat string
  temp string
  time string
  pitch string
  roll string
  yaw string
}

func BlankVitals() Vitals {
  return Vitals{
    bat: "--",
    temp: "--",
    time: "--",
    pitch: "--",
    roll: "--",
    yaw: "--",
  }
}

func VitalsFromTelementry(telemetry tello.Telemetry) Vitals {
  return Vitals {
    bat: telemetry.Bat + "%",
    temp: fmt.Sprintf("%s-%s°C", telemetry.Temp.Low, telemetry.Temp.High),
    time: telemetry.Time,
    pitch: telemetry.Pitch + "°",
    roll: telemetry.Roll + "°",
    yaw: telemetry.Yaw + "°",
  }
}
