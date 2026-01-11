package main

import (
	"math/rand/v2"
	"strconv"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var ansiColors = map[string]lipgloss.Color{
    "black": lipgloss.Color("0"),
    "red": lipgloss.Color("1"),
    "green": lipgloss.Color("2"),
    "yellow": lipgloss.Color("3"),
    "blue": lipgloss.Color("4"),
    "magenta": lipgloss.Color("5"),
    "cyan": lipgloss.Color("6"),
    "white": lipgloss.Color("7"),
    "grey": lipgloss.Color("8"),
    "bright-black": lipgloss.Color("8"),
    "bright-red": lipgloss.Color("9"),
    "bright-green": lipgloss.Color("10"),
    "bright-yellow": lipgloss.Color("11"),
    "bright-blue": lipgloss.Color("12"),
    "bright-magenta": lipgloss.Color("13"),
    "bright-cyan": lipgloss.Color("14"),
    "bright-white": lipgloss.Color("15"),
}

type Model struct {
	currentColor string
	vw int
	vh int
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
		
		case tea.KeyMsg:
			switch msg.String() {
				case "q":
					return m, tea.Quit
				case "r":
					m.currentColor = "red"
				case "g":
					m.currentColor = "green"
				case "b":
					m.currentColor = "blue"
				case "enter":
					m.currentColor = strconv.Itoa(rand.IntN(15))
			}
		
		case tea.WindowSizeMsg:
			m.vw = msg.Width
			m.vh = msg.Height
	}

	return m, nil
}

func (m Model) View() string {
	keyMap := map[string]string {
		"r": "Red",
		"g": "Green",
		"b": "Blue",
		"enter": "Rand",
		"q": "Quit",
	}

	var bg lipgloss.Color;

	if colour, ok := ansiColors[m.currentColor]; ok {
		bg = colour
	} else {
		bg = lipgloss.Color(m.currentColor)
	}

	keyMapStrs := []string{}

	for _, k := range []string{"r", "g", "b", "enter", "q"} {
		keyMapStrs = append(keyMapStrs, lipgloss.NewStyle().Bold(true).Render(k) + lipgloss.NewStyle().Background(bg).Render(" " + keyMap[k]))
	}

	return lipgloss.NewStyle().
		Width(m.vw).
		Height(m.vh).
		Padding(5).
		Background(bg).
		Render(
			lipgloss.JoinVertical(lipgloss.Left, keyMapStrs...),
		)
}

func main() {
	m := Model{currentColor: "black"}
	p := tea.NewProgram(m, tea.WithAltScreen())
	p.Run()
}
