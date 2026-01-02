package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/andrewhaine/go-tello/pkg/tello"
)

// This is where we'll re-create the basic CLI we had in v1
func main() {
  drone := tello.NewDrone()

  err := drone.ConnectDefault()

  if err != nil {
    fmt.Printf("Could not connect to drone: %s", err)
    os.Exit(1)
  }

  defer drone.CloseConnection()

  fmt.Println("Entering SDK mode...")
	drone.SendRawCmdString("command")
	
  fmt.Println("Ready for commands!")

  go printTelemetry(&drone)

  sendCommandsFromStdin(&drone)
}

func sendCommandsFromStdin(drone *tello.Drone) {
  for {
    reader := bufio.NewReader(os.Stdin)
    fmt.Print(">> ")
    
    text, _ := reader.ReadString('\n')
    cmdString := strings.TrimRight(text, "\n")

    if strings.TrimSpace(string(text)) == "emergency" {
      fmt.Println("Emergency stop!")
      drone.SendRawCmdString(cmdString)
      drone.CloseConnection()
			return
		}

    drone.SendRawCmdString(cmdString)
  }
}

func printTelemetry(drone *tello.Drone) {
  teleChan, err := drone.StreamTelemetry()

  if err != nil {
    fmt.Println("Could not stream telemetry: " + err.Error())
    return
  }

  for teleVal := range teleChan {
    s := fmt.Sprintf(
      "Tele: Battery: %s%%, Pitch: %s°, Yaw: %s°, Roll: %s°, Temp: %s°C - %s°C, Baro: %s, Time: %s, Altitude: %s, Height: %s",
      teleVal.Bat,
      teleVal.Pitch,
      teleVal.Yaw,
      teleVal.Roll,
      teleVal.Temp.Low,
      teleVal.Temp.High,
      teleVal.Baro,
      teleVal.Time,
      teleVal.Altitude,
      teleVal.Height,
    )
    fmt.Println(s)
  }
}
