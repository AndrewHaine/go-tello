package tui

import (
	"os/exec"
	"time"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type Dimensions struct {
  w int
  h int
}

type VitalsData struct {
  lastRec time.Time
  data Vitals
}

type TelloTui struct {
  title string
  connected bool
  dimensions Dimensions
  activeScreen Screen
  commandInput textinput.Model
  cmdChan chan<- string
  logMsgs []LogMessage
  logMsgChan <-chan LogMessage
  vitals VitalsData
  vitalsChan <-chan Vitals
  videoStreaming bool
  videoPlayerCmd *exec.Cmd
}

func NewModel() TelloTui {
  commandTextInput := textinput.New()
  commandTextInput.Placeholder = "e.g. `battery?`"
  commandTextInput.CharLimit = 25
  commandTextInput.Width = 20

  return TelloTui {
    title: "GO TELLO TUI",
    connected: false,
    dimensions: Dimensions{w: 0, h: 0},
    activeScreen: SCREEN_MAIN,
    commandInput: commandTextInput,
    vitals: VitalsData{data: BlankVitals()},
  }
}

func (tt *TelloTui) SetCmdChan(cmdChan chan string) {
  tt.cmdChan = cmdChan
}

func (tt *TelloTui) SetLogMsgChan(logMsgChan chan LogMessage) {
  tt.logMsgChan = logMsgChan
}

func (tt *TelloTui) SetVitalsChan(vitalsChan chan Vitals) {
  tt.vitalsChan = vitalsChan
}

type CheckConnectionMsg time.Time

type VitalsMsg struct {
  Vitals Vitals
}

type LogMsgMsg struct {
  LogMsg LogMessage
}

func CheckConnection() tea.Cmd {
  return tea.Tick(time.Second, func (currTime time.Time) tea.Msg {
    return CheckConnectionMsg(currTime)
  })
}

func ListenForDroneMsg(tt TelloTui) tea.Cmd {
  return func() tea.Msg {
    select {
      case vitals := <- tt.vitalsChan:
        return VitalsMsg{Vitals: vitals}
      case logMsg := <- tt.logMsgChan:
        return LogMsgMsg{LogMsg: logMsg}
    }
  }
}

func (tt TelloTui) Init() tea.Cmd {
  return tea.Batch(tea.EnterAltScreen, ListenForDroneMsg(tt), CheckConnection())
}

func (tt TelloTui) Cleanup() error {
  return nil
}
