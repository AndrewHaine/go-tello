package tui

import (
	"slices"
	"time"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
)

var (
  foregroundColour = lipgloss.Color("#90FA7B")
  secondaryForegroundColour = lipgloss.Color("#6DA063")

  maxWidth = 74
  maxCaptainsLogMessages = 5

  baseStyle = lipgloss.NewStyle().Foreground(foregroundColour)
  
  textStyle = baseStyle
  secondaryTextStyle = lipgloss.NewStyle().Foreground(secondaryForegroundColour)
  boldTextStyle = textStyle.Bold(true)
)

func boldNormalTextPairString(bold string, normal string, useSecondaryForeground bool) string {
  if useSecondaryForeground {
    return boldTextStyle.Render(bold) + secondaryTextStyle.Render(normal)
  }

  return boldTextStyle.Render(bold) + textStyle.Render(normal)
}

func renderHeader(tt TelloTui) string {
  headerTable := table.New().
    Width(maxWidth).
    Border(lipgloss.NormalBorder()).
    BorderStyle(baseStyle).
    StyleFunc(func(row, col int) lipgloss.Style {
      return baseStyle.Align(lipgloss.Center)
    })

  connectionString := "Not connected";

  if tt.connected {
    connectionString = "Connected" 
  }

  headerTable.Row(
    boldTextStyle.Render(tt.title),
    textStyle.Render("Status ") + boldTextStyle.Render(connectionString),
    boldNormalTextPairString("q", " quit", true),
  )

  return headerTable.String()
}

func renderFlightCtrl() string {
  flightCtrlString := lipgloss.NewStyle().
    Border(lipgloss.NormalBorder()).
    BorderForeground(foregroundColour).
    Padding(0, 1).
    MarginRight(2).
    Width((maxWidth / 2) - 6).
    Height(15)
  
  flightCtrlRows := [][]string{
    {boldNormalTextPairString("w", " forwards", false), boldNormalTextPairString("↑", " up", false)},
    {boldNormalTextPairString("a", " left", false), boldNormalTextPairString("↓", " down", false)},
    {boldNormalTextPairString("s", " backwards", false), boldNormalTextPairString("→", " right (yaw)", false)},
    {boldNormalTextPairString("d", " right", false), boldNormalTextPairString("←", " left (yaw)", false)},
    {boldNormalTextPairString("t", " take off", false)},
    {boldNormalTextPairString("l", " land", false)},
    {boldNormalTextPairString("e", " emergency", false)},
  }

  flightCtrlTable := table.New().
    Border(lipgloss.HiddenBorder()).
    StyleFunc(func(row, col int) lipgloss.Style {
      if col == 1 {
        return baseStyle.Padding(0, 0)
      }

      return baseStyle.Padding(0, 2, 0, 0)
    }).
    Rows(flightCtrlRows...)

  flighCtrlContents := lipgloss.JoinVertical(
    lipgloss.Left,
    boldTextStyle.Render("<flight controls>"),
    flightCtrlTable.String(),
  )

  return flightCtrlString.Render(flighCtrlContents)
}

func renderVitals(tt TelloTui) string {
  vitalsBlock := lipgloss.NewStyle().
    Border(lipgloss.NormalBorder()).
    BorderForeground(foregroundColour).
    Padding(0, 1).
    Width((maxWidth / 2) + 2)

  vitals := tt.vitals.data
  
  flightCtrlRows := [][]string{
    {boldNormalTextPairString("bat ", vitals.bat, false), boldNormalTextPairString("pitch ", vitals.pitch, false)},
    {boldNormalTextPairString("temp ", vitals.temp, false), boldNormalTextPairString("roll ", vitals.roll, false)},
    {boldNormalTextPairString("height ", vitals.height, false), boldNormalTextPairString("yaw ", vitals.yaw, false)},

  }

  vitalsTable := table.New().
    Border(lipgloss.HiddenBorder()).
    StyleFunc(func(row, col int) lipgloss.Style {
      if col == 1 {
        return baseStyle.Padding(0, 0)
      }

      return baseStyle.Padding(0, 4, 0, 0)
    }).
    Rows(flightCtrlRows...)

  vitalsBlockContents := lipgloss.JoinVertical(
    lipgloss.Left,
    boldTextStyle.Render("<vitals>"),
    vitalsTable.String(),
  )

  return vitalsBlock.Render(vitalsBlockContents)
}

func renderCaptainsLog(tt TelloTui) string {
  captainsLogBlock := lipgloss.NewStyle().
    Border(lipgloss.NormalBorder()).
    BorderForeground(foregroundColour).
    Padding(0, 1).
    Width((maxWidth / 2) + 2)

  rows := []string{};
  
  msgs := tt.logMsgs
  // Explicitly sort the messages by time
  slices.SortFunc(msgs, func(a, b LogMessage) int {
    return a.Time.Compare(b.Time)
  })

  msgCount := len(msgs)

  if msgCount < 1 {
    rows = append(rows, textStyle.Render("Empty"))
  } else {
    slices.Reverse(tt.logMsgs)

    for i, msg := range tt.logMsgs {
      if i > maxCaptainsLogMessages {
        break
      }
      row := textStyle.Render(msg.Time.Format(time.TimeOnly)) + " " + boldTextStyle.Render(msg.Message)
      rows = append(rows, row)
    }
  }

  slices.Reverse(rows)

  // Pad the remaining rows
  if (msgCount < maxCaptainsLogMessages) {
    toPad := maxCaptainsLogMessages - msgCount
    for range toPad {
      rows = append(rows, textStyle.Render("   "))
    }
  }

  rows = append([]string{boldTextStyle.Render("<captain's log>")}, rows...)

  captainsLogContents := lipgloss.JoinVertical(
    lipgloss.Left,
    rows...
  )

  return captainsLogBlock.Render(captainsLogContents)
}

func renderFooter(tt TelloTui) string {
  footer := lipgloss.NewStyle().
    AlignHorizontal(lipgloss.Center).
    Width(maxWidth).
    Border(lipgloss.NormalBorder()).
    BorderForeground(foregroundColour).
    Inherit(baseStyle)

  var commands []string

  videoLabel := "start video"

  if tt.videoStreaming {
    videoLabel = "stop video"
  }
  
  switch tt.activeScreen {
    case SCREEN_COMMAND:
      commands = []string{
        boldNormalTextPairString("esc", " back", true),
        textStyle.Render("   "),
        boldNormalTextPairString("enter", " submit", true),
      }
    case SCREEN_MAIN:
      commands = []string{
        boldNormalTextPairString("c", " command", true),
        textStyle.Render("   "),
        boldNormalTextPairString("v ", videoLabel, true),
      }
  }
  
  return footer.Render(lipgloss.JoinHorizontal(
    lipgloss.Top,
    commands...
  ))
}

func renderMainScreen(tt TelloTui) string {
  return lipgloss.JoinHorizontal(
      lipgloss.Top,
      renderFlightCtrl(),
      lipgloss.JoinVertical(
        lipgloss.Left,
        renderVitals(tt),
        renderCaptainsLog(tt),
      ),
    )
}

func renderCommandScreen(tt TelloTui) string {
  commandContainer := lipgloss.NewStyle().Height(17).Inherit(baseStyle)

  return commandContainer.Render(
    lipgloss.JoinVertical(
      lipgloss.Center,
      boldTextStyle.MarginTop(1).Render("Enter your command:"),
      tt.commandInput.View(),   
    ),
  )
}

func (tt TelloTui) View() string {
  scene := baseStyle.
    Width(tt.dimensions.w).
    Height(tt.dimensions.h).
    Padding(1, 2).
    Align(lipgloss.Center)

  var screenString string

  switch tt.activeScreen {
    case SCREEN_COMMAND:
      screenString = renderCommandScreen(tt)
    case SCREEN_MAIN:
      screenString = renderMainScreen(tt)
    default:
      screenString = renderMainScreen(tt)
  }

  return scene.Render(
    lipgloss.JoinVertical(
      lipgloss.Center,
      renderHeader(tt),
      screenString,
      renderFooter(tt),
    ),
  )
}
