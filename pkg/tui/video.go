package tui

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
)

const (
  VIDEO_ADDR = "0.0.0.0:11111"
)

func (tt *TelloTui) StartVideo() (error) {
  if tt.videoStreaming {
    return errors.New("Video player already started")
  }

  cmd := exec.Command("ffplay",
    "-f", "h264",
    "-window_title", "TelloVision",
    fmt.Sprintf("udp://%s", VIDEO_ADDR),
  )

  err := cmd.Start()

  if err != nil {
    return err
  }

  tt.videoStreaming = true
  tt.videoPlayerCmd = cmd

  return nil
}

func (tt *TelloTui) StopVideo() (error) {
  if !tt.videoStreaming {
    return nil
  }

  err := tt.videoPlayerCmd.Process.Signal(os.Interrupt)

  if err != nil {
    return err
  }

  return nil
}
