package tui

import (
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
  vitals VitalsData
  vitalsChan chan Vitals
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

func (tt *TelloTui) SetVitalsChan(vitalsChan chan Vitals) {
  tt.vitalsChan = vitalsChan
}

type VitalsMsg struct {
  Vitals Vitals
}

func ListenForTelemetry(tt TelloTui) tea.Cmd {
  return func() tea.Msg {
    vitals := <- tt.vitalsChan
    return VitalsMsg{Vitals: vitals}
  }
}

func (tt TelloTui) Init() tea.Cmd {
  return tea.Batch(tea.EnterAltScreen, ListenForTelemetry(tt))
}
