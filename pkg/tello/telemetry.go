package tello

import (
	"errors"
	"strings"
)

type TemperatureTelemetryValue struct {
  Low string
  High string
}

type MultiDirectionalTelemetryValue struct {
  X string
  Y string
  Z string
}

type Telemetry struct {
  Pitch string
  Roll string
  Yaw string
  Altitude string
  Height string
  Time string
  Bat string
  Baro string
  Temp TemperatureTelemetryValue
  Speed MultiDirectionalTelemetryValue
  Acceleration MultiDirectionalTelemetryValue
}

func (drone *Drone) StreamTelemetry() (<-chan Telemetry, error) {
  if !drone.telemetryConnState.connected {
    return nil, errors.New("Telemetry connection not established")
  }

  teleChan := make(chan Telemetry)

  go func() {
    for {
      teleBuff := make([]byte, 4096)
      drone.telemetryConn.Read(teleBuff)
      teleChan <- telemetryFromBuff(teleBuff)
    }
  }()

  return teleChan, nil
}

func telemetryFromBuff(buff []byte) Telemetry {
  valuePairs := strings.Split(string(buff), ";")
  valueMap := map[string]string{}

  for _, vp := range valuePairs {
    parts := strings.Split(vp, ":")
    if (len(parts) != 2) {
      continue
    }
    valueMap[parts[0]] = parts[1]
  }

  tempValue := TemperatureTelemetryValue{
    High: valueMap["temph"],
    Low: valueMap["templ"],
  }

  speedVal := MultiDirectionalTelemetryValue{
    X: valueMap["vgx"],
    Y: valueMap["vgy"],
    Z: valueMap["vgz"],
  }

  accelerationVal := MultiDirectionalTelemetryValue{
    X: valueMap["agx"],
    Y: valueMap["agy"],
    Z: valueMap["agz"],
  }

  return Telemetry{
    Pitch: valueMap["pitch"],
    Roll: valueMap["roll"],
    Yaw: valueMap["yaw"],
    Altitude: valueMap["tof"],
    Height: valueMap["h"],
    Bat: valueMap["bat"],
    Baro: valueMap["baro"],
    Time: valueMap["time"],
    Temp: tempValue,
    Speed: speedVal,
    Acceleration: accelerationVal,
  }
}
