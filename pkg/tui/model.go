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

func (tt TelloTui) Init() tea.Cmd {
  return nil
}
