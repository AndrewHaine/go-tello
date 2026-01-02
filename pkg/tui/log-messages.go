package tui

import (
	"time"

	"github.com/andrewhaine/go-tello/pkg/tello"
)

type LogMessage struct {
  Time time.Time
  Message string
}

func LogMsgFromTelloMsg(telloMsg tello.Message) LogMessage {
  return LogMessage{Time: telloMsg.Time, Message: telloMsg.Message}
}

func (tt *TelloTui) AppendLogMsg(logMsg LogMessage) {
  if (len(logMsg.Message) < 1) {
    return
  }

  tt.logMsgs = append(tt.logMsgs, logMsg)
}
