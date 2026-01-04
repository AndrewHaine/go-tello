package tui

import (
	"log"
	"math"
	"time"
)

func CheckConnectionFromVitals(tt *TelloTui) {
  sinceLastVitals := time.Since(tt.vitals.lastRec)
  sinceLastVitalsSecs := sinceLastVitals.Seconds()

  if tt.connected && sinceLastVitalsSecs > 5 {
    log.Printf("Connection Lost: Vitals last received %.0f seconds ago", math.Ceil(sinceLastVitalsSecs))
    tt.logMsgChan <- LogMessage{Time: time.Now(), Message: "Connection Lost", Type: LOG_DEBUG}
    tt.connected = false
    return
  }

  if !tt.connected && sinceLastVitalsSecs < 5 {
    log.Printf("Connection Established: Vitals last received %.0f seconds ago", math.Ceil(sinceLastVitalsSecs))
    tt.logMsgChan <- LogMessage{Time: time.Now(), Message: "Connection Established", Type: LOG_DEBUG}
    tt.connected = true
  }
}
