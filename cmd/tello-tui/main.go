package main

import (
	"fmt"
	"os"

	"github.com/andrewhaine/go-tello/pkg/tui"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
  p := tea.NewProgram(
    tui.NewModel(),
    tea.WithAltScreen(),
  )
  
  if _, err := p.Run(); err != nil {
    fmt.Printf("Error occurred starting program: %v", err)
    os.Exit(1)
  }
}
