package tello

import (
	"fmt"
	"net"
)

const (
	defaultVideoAddr = "0.0.0.0"
  defaultVideoPort = "11111"
)

func (drone *Drone) StreamVideo() (*net.UDPConn, error) {
	if drone.streamingVideo {
		return drone.videoConn, nil
	}

	drone.cmdConn.Write([]byte("streamon"))

	remoteUdpAddr, err := net.ResolveUDPAddr("udp", fmt.Sprintf("%s:%s", defaultVideoAddr, defaultVideoPort))

	if err != nil {
		return nil, err
	}

	conn, err := net.ListenUDP("udp", remoteUdpAddr)

	if err != nil {
		return nil, err
	}

	drone.streamingVideo = true
	drone.videoConn = conn
	return conn, nil
}
