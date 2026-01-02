package tui

import (
	"log"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

func updateMainScreenKeyMsg (tt TelloTui, msg tea.KeyMsg) (tea.Model, tea.Cmd) {
  log.Println("Key: " + msg.String())

  switch msg.String() {
    case "q", "ctrl+c":
      return tt, tea.Quit
    case "c":
      tt.activeScreen = SCREEN_COMMAND
      tt.commandInput.Reset()
      tt.commandInput.Focus()
    case "t":
      tt.cmdChan <- "takeoff"
    case "l":
      tt.cmdChan <- "land"
    case "e":
      tt.cmdChan <- "emergency"
    case "w":
      tt.cmdChan <- "forward 50"
    case "a":
      tt.cmdChan <- "left 50"
    case "s":
      tt.cmdChan <- "back 50"
    case "d":
      tt.cmdChan <- "right 50"
    case "up":
      tt.cmdChan <- "up 50"
    case "down":
      tt.cmdChan <- "down 50"
    case "left":
      tt.cmdChan <- "ccw 60"
    case "right":
      tt.cmdChan <- "cw 60"
    
  }

  return tt, nil
}

func updateCommandScreenKeyMsg (tt TelloTui, msg tea.KeyMsg) (tea.Model, tea.Cmd) {
  var cmd tea.Cmd

  log.Println("Key: " + msg.String())

  switch (msg.String()) {
    case "q", "ctrl+c":
      return tt, tea.Quit
    case "esc":
      tt.activeScreen = SCREEN_MAIN
      tt.commandInput.Blur()
    case "enter":
      tt.cmdChan <- tt.commandInput.Value()
      tt.commandInput.Blur()
      tt.activeScreen = SCREEN_MAIN
  }

  tt.commandInput, cmd = tt.commandInput.Update(msg)
  return tt, cmd;
}

func (tt TelloTui) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
  var cmd tea.Cmd

  switch msg := msg.(type) {
    case tea.KeyMsg:
      switch tt.activeScreen {
        case SCREEN_MAIN:
          return updateMainScreenKeyMsg(tt, msg)
        case SCREEN_COMMAND:
          return updateCommandScreenKeyMsg(tt, msg)
      }
    
    case VitalsMsg:
      tt.vitals = VitalsData{ data: msg.Vitals, lastRec: time.Now() }
      return tt, ListenForDroneMsg(tt)

    case LogMsgMsg:
      tt.AppendLogMsg(msg.LogMsg)
      return tt, ListenForDroneMsg(tt)

    case tea.WindowSizeMsg:
      tt.dimensions.w = msg.Width
      tt.dimensions.h = msg.Height
  }

  return tt, cmd
}
