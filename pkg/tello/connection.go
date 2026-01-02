package tello

import (
	"errors"
	"fmt"
	"net"
)

const (
  defaultCmdAddr = "192.168.10.1"
  defaultCmdPort = "8889"
  defaultHostCmdPort = "8500"
  defaultTelemetryAddr = "0.0.0.0"
  defaultTelemetryPort = "8890"
  defaultHostTelemetryPort = "8501"
)

func NewConnection(addr string, port string, hostPort string) Connection {
  return Connection{addr: addr, port: port}
}

func DefaultCmdConnection() Connection {
  return NewConnection(defaultCmdAddr, defaultCmdPort, defaultHostCmdPort)
}

func DefaultTelemetryConnection() Connection {
  return NewConnection(defaultTelemetryAddr, defaultTelemetryPort, defaultHostTelemetryPort)
}

func makeUDPConn(connDetails Connection, connState *ConnectionState, isServer bool) (*net.UDPConn, error) {
if (connState.connected) {
    return nil, errors.New("Connection already in place")
  }

  if connState.connecting {
    return nil, errors.New("Connection already in progress")
  }

  remoteUdpAddr, err := net.ResolveUDPAddr("udp", fmt.Sprintf("%s:%s", connDetails.addr, connDetails.port))
  
  connState.connecting = true
  
  if err != nil {
    connState.connecting = false
    return nil, err
  }

  var conn *net.UDPConn
  var connErr error

  if isServer {
    // Create a local UDP addr so that the drone has a consistent place to send responses
    localUdpAddr, err := net.ResolveUDPAddr("udp", ":" + connDetails.hostPort)

    if err != nil {
      connState.connecting = false
      return nil, err
    }

    conn, connErr = net.DialUDP("udp", localUdpAddr, remoteUdpAddr)
  } else {
    conn, connErr = net.ListenUDP("udp", remoteUdpAddr)
  }

  if (connErr != nil) {
    connState.connecting = false
    return nil, err
  }

  connState.connecting = false
  connState.connected = true

  return conn, nil
}

func (drone *Drone) Connect(cmdConn Connection, telemetryConn Connection) (err error) {
  cmdConnection, err := makeUDPConn(cmdConn, &drone.cmdConnState, true)

  if (err != nil) {
    return errors.New("Could not make command connection: " + err.Error())
  }

  telemetryConnection, err := makeUDPConn(telemetryConn, &drone.telemetryConnState, false)

  if (err != nil) {
    return errors.New("Could not make telemetry connection: " + err.Error())
  }

  drone.cmdConn = cmdConnection
  drone.telemetryConn = telemetryConnection

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
