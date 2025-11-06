package main

import (
	"bufio"
	"fmt"
	"os"
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

  go printCmdMessages(&drone)
	
  fmt.Println("Ready for commands!")
  sendCommandsFromStdin(&drone)
}

func sendCommandsFromStdin(drone *tello.Drone) {
  for {
    reader := bufio.NewReader(os.Stdin)
    fmt.Print(">> ")
    
    text, _ := reader.ReadString('\n')
    cmdString := strings.TrimRight(text, "\n")

    drone.SendRawCmdString([]byte(cmdString))

    if strings.TrimSpace(string(text)) == "STOP" {
			fmt.Println("Stopping")
			return
		}
  }
}

func printCmdMessages(drone *tello.Drone) {
	for {
    message := <-drone.CmdResponseChan
    fmt.Printf("🚁 Message from Tello drone: %s\n>>", message)
  }
}
