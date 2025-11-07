package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"

	"github.com/andrewhaine/go-tello/tello"
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
	drone.SendRawCmdString([]byte("command"))
	
  fmt.Println("Ready for commands!")
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
      drone.SendRawCmdString([]byte(cmdString))
      drone.CloseConnection()
			return
		}

    if strings.TrimSpace(string(text)) == "M" {
      messageStrings := []string{}
      for _, droneMsg := range drone.GetMessages() {
        messageStrings = append(messageStrings, droneMsg.Message)
      }
      fmt.Println(slices.Concat(messageStrings))
      continue
    }

    drone.SendRawCmdString([]byte(cmdString))
  }
}
