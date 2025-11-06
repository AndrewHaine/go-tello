package tello

import (
	"errors"
	"fmt"
	"net"
)

const (
  defaultCmdAddr = "192.168.10.1"
  defaultCmdPort = "8889"
  defaultTelemetryAddr = "0.0.0.0"
  defaultTelemetryPort = "8890"
)

type Connection struct {
  addr string
  port string
}

type ConnectionState struct {
  connected bool
  connecting bool
}

type Drone struct {
  cmdConn *net.UDPConn
  cmdConnState ConnectionState
  telemetryConn *net.UDPConn
  telemetryConnState ConnectionState
  CmdResponseChan chan []byte
  // telemetryChan chan State
}

func NewDrone() Drone {
  drone := Drone{}
  return drone
}

func NewConnection(addr string, port string) Connection {
  return Connection{addr: addr, port: port}
}

func DefaultCmdConnection() Connection {
  return NewConnection(defaultCmdAddr, defaultCmdPort)
}

func DefaultTelemetryConnection() Connection {
  return NewConnection(defaultTelemetryAddr, defaultTelemetryPort)
}

func makeUDPConn(connDetails Connection, connState *ConnectionState) (conn *net.UDPConn, err error) {
if (connState.connected) {
    return nil, errors.New("Connection already in place")
  }

  if connState.connecting {
    return nil, errors.New("Connection already in progress")
  }

  udpAddr, err := net.ResolveUDPAddr("udp", fmt.Sprintf("%s:%s", connDetails.addr, connDetails.port))
  
  connState.connecting = true
  
  if err != nil {
    connState.connecting = false
    return nil, err
  }

  cmdConn, err := net.DialUDP("udp", nil, udpAddr)

  if (err != nil) {
    connState.connecting = false
    return nil, err
  }

  connState.connecting = false
  connState.connected = true

  return cmdConn, nil
}

func (drone *Drone) Connect(cmdConn Connection, telemetryConn Connection) (err error) {
  cmdConnection, err := makeUDPConn(cmdConn, &drone.cmdConnState)

  if (err != nil) {
    return errors.New("Could not make command connection: " + err.Error())
  }

  telemetryConnection, err := makeUDPConn(telemetryConn, &drone.telemetryConnState)

  if (err != nil) {
    return errors.New("Could not make telemetry connection: " + err.Error())
  }

  drone.cmdConn = cmdConnection
  drone.telemetryConn = telemetryConnection

  go drone.startCmdResponseListener()

  return nil
}

func (drone *Drone) CloseConnection() {
  drone.cmdConn.Close()
  drone.cmdConnState.connected = false
  drone.telemetryConn.Close()
  drone.telemetryConnState.connected = false
}

func (drone *Drone) ConnectDefault() (err error) {
  return drone.Connect(DefaultCmdConnection(), DefaultTelemetryConnection())
}

func (drone *Drone) startCmdResponseListener() {
  drone.CmdResponseChan = make(chan []byte)

  for {
    messageBuff := make([]byte, 4096)
    drone.cmdConn.Read(messageBuff)
    drone.CmdResponseChan <- messageBuff
  }
}

func (drone *Drone) SendRawCmdString(cmdString []byte) {
  drone.cmdConn.Write(cmdString);
}
