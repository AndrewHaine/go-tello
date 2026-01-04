package tui

import (
	"time"

	"github.com/andrewhaine/go-tello/pkg/tello"
)

type LogMessageType int

const (
  LOG_INFO LogMessageType = iota
  LOG_DEBUG
  LOG_ERROR
)

type LogMessage struct {
  Time time.Time
  Message string
  Type LogMessageType
}

func LogMsgFromTelloMsg(telloMsg tello.Message) LogMessage {
  return LogMessage{Time: telloMsg.Time, Message: telloMsg.Message, Type: LOG_INFO}
}

func (tt *TelloTui) AppendLogMsg(logMsg LogMessage) {
  if (len(logMsg.Message) < 1) {
    return
  }

  if (len(logMsg.Message) > 30) {
    logMsg.Message = logMsg.Message[:29]
  }

  tt.logMsgs = append(tt.logMsgs, logMsg)
}
