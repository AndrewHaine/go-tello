package tello

import "time"

type Message struct {
  Time time.Time
  Message string
}

func (drone *Drone) startCmdMessageLogger() {
  for cmdMessage := range drone.cmdResponseChan  {
    drone.cmdMessagesMu.Lock()
    logMsg := Message{Time: time.Now(), Message: cmdMessage}
    drone.cmdMessages = append(drone.cmdMessages, logMsg)
    drone.cmdMessagesMu.Unlock()
  }
}

func (drone *Drone) GetMessages() []Message {
  drone.cmdMessagesMu.Lock()
  defer drone.cmdMessagesMu.Unlock()
  return drone.cmdMessages
}
