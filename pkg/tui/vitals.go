package tui

import (
	"fmt"

	"github.com/andrewhaine/go-tello/pkg/tello"
)

type Vitals struct {
  bat string
  temp string
  height string
  pitch string
  roll string
  yaw string
}

func BlankVitals() Vitals {
  return Vitals{
    bat: "--",
    temp: "--",
    height: "--",
    pitch: "--",
    roll: "--",
    yaw: "--",
  }
}

func (vitals Vitals) String() string {
  return fmt.Sprintf("Bat: %s;Temp: %s;Height: %s; Pitch: %s; Roll: %s; Yaw: %s", vitals.bat, vitals.temp, vitals.height, vitals.pitch, vitals.roll, vitals.yaw)
}

func VitalsFromTelementry(telemetry tello.Telemetry) Vitals {
  return Vitals {
    bat: telemetry.Bat + "%",
    temp: fmt.Sprintf("%s-%s°C", telemetry.Temp.Low, telemetry.Temp.High),
    height: telemetry.Height + "cm",
    pitch: telemetry.Pitch + "°",
    roll: telemetry.Roll + "°",
    yaw: telemetry.Yaw + "°",
  }
}
