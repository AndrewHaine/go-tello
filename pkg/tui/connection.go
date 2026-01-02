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
    tt.connected = false
    return
  }

  if !tt.connected && sinceLastVitalsSecs < 5 {
    log.Printf("Connection Re-establoshed: Vitals last received %.0f seconds ago", math.Ceil(sinceLastVitalsSecs))
    tt.connected = true
  }
}
