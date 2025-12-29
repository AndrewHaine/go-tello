package tello

import (
	"net"
	"sync"
)

type Connection struct {
  addr string
  port string
  hostPort string
}

type ConnectionState struct {
  connected bool
  connecting bool
}

type Drone struct {
  cmdConn *net.UDPConn
  cmdConnState ConnectionState
  cmdMessagesMu sync.Mutex
  cmdMessages []Message
  cmdResponseChan chan string
  telemetryConn *net.UDPConn
  telemetryConnState ConnectionState
}

func NewDrone() Drone {
  return Drone{
    cmdResponseChan: make(chan string, 20),
  }
}

func (drone *Drone) SendRawCmdString(cmdString string) {
  drone.cmdConn.Write([]byte(cmdString));
}
