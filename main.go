package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

const (
	controlAddr = "192.168.10.1"
	controlPort = "8889"
	telemetryAddr = "0.0.0.0"
	telemetryPort = "8890"
	videoPort = "11111"
)

func main() {
	controlLocation := fmt.Sprintf("%s:%s", controlAddr, controlPort)
	udpAddr, err := net.ResolveUDPAddr("udp", controlLocation)

	if err != nil {
		fmt.Printf("Could not resolve drone address: %s", err)
		os.Exit(1)
	}

	conn, err := net.DialUDP("udp", nil, udpAddr)

	if err != nil {
		fmt.Printf("Could not connect to drone: %s", err)
		os.Exit(1)
	}

	defer conn.Close()
	
	fmt.Println("Entering SDK mode...")
	conn.Write([]byte("command"))

	fmt.Println("Ready for commands!")

	messageChan := receiveMessagesFromTello(conn)

	go printMessages(messageChan)

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print(">> ")
		text, _ := reader.ReadString('\n')

		command := strings.TrimRight(text, "\n")
		
		conn.Write([]byte(command))

		if strings.TrimSpace(string(text)) == "STOP" {
			fmt.Println("Stopping")
			return
		}
	}
}

func receiveMessagesFromTello(conn *net.UDPConn) chan []byte { 
	telloMessageChan := make(chan []byte)

	go func() {
		for {
			message := make([]byte, 256)
			conn.ReadFromUDP(message)
			telloMessageChan <- message
		}
	}()

	return telloMessageChan
}

func printMessages(messageChan chan []byte) {
	for {
			message := <-messageChan
			fmt.Printf("🚁 Message from Tello: %s\n>>", message)
		}
}
