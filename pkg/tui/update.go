package tui

import (
	tea "github.com/charmbracelet/bubbletea"
)

func updateMainScreenKeyMsg (tt TelloTui, msg tea.KeyMsg) (tea.Model, tea.Cmd) {
  switch msg.String() {
    case "q", "ctrl+c":
      return tt, tea.Quit
    case "c":
      tt.activeScreen = SCREEN_COMMAND
      tt.commandInput.Reset()
      tt.commandInput.Focus()
  }

  return tt, nil
}

func updateCommandScreenKeyMsg (tt TelloTui, msg tea.KeyMsg) (tea.Model, tea.Cmd) {
  var cmd tea.Cmd

  switch (msg.String()) {
    case "q", "ctrl+c":
      return tt, tea.Quit
    case "esc":
      tt.activeScreen = SCREEN_MAIN
      tt.commandInput.Blur()
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
    
    case tea.WindowSizeMsg:
      tt.dimensions.w = msg.Width
      tt.dimensions.h = msg.Height
  }

  return tt, cmd
}
