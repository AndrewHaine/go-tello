package main

import (
	"fmt"
	"os"

	"github.com/andrewhaine/go-tello/pkg/tello"
	"github.com/andrewhaine/go-tello/pkg/tui"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
  drone := tello.NewDrone();

  err := drone.ConnectDefault()

  if err != nil {
    fmt.Printf("Could not connect to drone: %s", err)
    os.Exit(1)
  }

  defer drone.CloseConnection()
  
  drone.SendRawCmdString("command")

  m := tui.NewModel()

  vitalsChan := make(chan tui.Vitals)
  m.SetVitalsChan(vitalsChan)

  go streamTelemetryToVitals(&drone, vitalsChan)

  p := tea.NewProgram(m)
  
  if _, err := p.Run(); err != nil {
    fmt.Printf("Error occurred starting program: %v", err)
    os.Exit(1)
  }
}

func streamTelemetryToVitals(drone *tello.Drone, vitalsChan chan tui.Vitals) {
  telemetryChan, err := drone.StreamTelemetry()

  if err != nil {
    fmt.Println("Could not stream telemetry " + err.Error())
    return
  }

  for telemetry := range telemetryChan {
    vitals := tui.VitalsFromTelementry(telemetry)
    vitalsChan <- vitals
  }
}
