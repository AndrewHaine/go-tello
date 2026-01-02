package tello

import (
	"errors"
	"log"
	"time"
)

type Message struct {
  Time time.Time
  Message string
}

func (drone *Drone) StreamMessages() (<-chan Message, error) {
  if !drone.cmdConnState.connected {
    return nil, errors.New("Command connection not established")
  }

  msgChan := make(chan Message)

  go func() {
    for {
      msgBuff := make([]byte, 1024)
      n, err := drone.cmdConn.Read(msgBuff)

      if err != nil {
        log.Println("Error reading from command connection ", err.Error())
      }

      msg := Message{Time: time.Now(), Message: string(msgBuff[:n])}
      msgChan <- msg
    }
  }()

  return msgChan, nil
}
