package tello

import (
	"net"
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
  messagesStreaming bool
  telemetryStreaming bool
  telemetryConn *net.UDPConn
  telemetryConnState ConnectionState
  streamingVideo bool
  videoConn *net.UDPConn
}

func NewDrone() Drone {
  return Drone{}
}

func (drone *Drone) SendRawCmdString(cmdString string) {
  drone.cmdConn.Write([]byte(cmdString));
}
