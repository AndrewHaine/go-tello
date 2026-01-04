package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/andrewhaine/go-tello/pkg/tello"
	"github.com/andrewhaine/go-tello/pkg/tui"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
  drone := tello.NewDrone();
  
  f, logErr := tea.LogToFile("debug.log", "debug")
	if logErr != nil {
		fmt.Println("fatal:", logErr)
		os.Exit(1)
	}
	defer f.Close()

  err := drone.ConnectDefault()

  if err != nil {
    fmt.Printf("Could not connect to drone: %s", err)
    os.Exit(1)
  }

  defer drone.CloseConnection()
  
  drone.SendRawCmdString("command")

  m := tui.NewModel()

  // Create a channel to recieve commands from the TUI model 
  cmdChan  := make(chan string)
  m.SetCmdChan(cmdChan)
  go listenForTuiCmd(&drone, cmdChan)

  // Create a channel to send messages to the TUI model
  logMsgChan := make(chan tui.LogMessage)
  m.SetLogMsgChan(logMsgChan)  
  go streamMessages(&drone, logMsgChan)

  // Create a channel to send telemetry data to the TIU model
  vitalsChan := make(chan tui.Vitals)
  m.SetVitalsChan(vitalsChan)
  go streamTelemetryToVitals(&drone, vitalsChan)

  p := tea.NewProgram(m)
  
  if _, err := p.Run(); err != nil {
    fmt.Printf("Error occurred starting program: %v", err)
    drone.SendRawCmdString("emergency")
    os.Exit(1)
  }
}

func listenForTuiCmd(drone *tello.Drone, cmdChan <-chan string) {
  for cmd := range cmdChan {
    drone.SendRawCmdString(cmd)
  }
}

func streamMessages(drone *tello.Drone, logMsgChan chan<- tui.LogMessage) {
  msgChan, err := drone.StreamMessages()

  if err != nil {
    log.Println("Could not stream messages " + err.Error())
    return
  }

  for msg := range msgChan {
    log.Println("Msg rec: " + strings.Trim(msg.Message, " "))
    logMsg := tui.LogMsgFromTelloMsg(msg)
    logMsgChan <- logMsg
  }
}

func streamTelemetryToVitals(drone *tello.Drone, vitalsChan chan<- tui.Vitals) {
  telemetryChan, err := drone.StreamTelemetry()

  if err != nil {
    log.Println("Could not stream telemetry " + err.Error())
    return
  }

  for telemetry := range telemetryChan {
    vitals := tui.VitalsFromTelementry(telemetry)
    vitalsChan <- vitals
  }
}
