package tui

type Screen int

const (
  SCREEN_MAIN Screen = iota
  SCREEN_COMMAND
)

var screenName = map[Screen]string {
  SCREEN_MAIN: "main",
  SCREEN_COMMAND: "command",
}

func (screen Screen) String() string {
  return screenName[screen]
}
