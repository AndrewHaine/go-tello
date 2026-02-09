package web

import (
	"bytes"
	"fmt"
	"os/exec"
	"path/filepath"
	"sync"
	"time"

	"github.com/andrewhaine/go-tello/pkg/tello"
	"github.com/pion/webrtc/v3"
	"github.com/pion/webrtc/v3/pkg/media"
	"github.com/pion/webrtc/v3/pkg/media/h264reader"
)

type VideoServer struct {
	Mu sync.RWMutex
	nalBuffer [][]byte
	maxBufferSize int
	screenshotDir string
	VideoTrack *webrtc.TrackLocalStaticSample
}

func NewVideoServer(screenshotDir string) VideoServer {
	videoTrack, _ := webrtc.NewTrackLocalStaticSample(
		webrtc.RTPCodecCapability{MimeType: webrtc.MimeTypeH264},
		"video",
		"stream",
	)

	return VideoServer{
		maxBufferSize: 75,
		screenshotDir: screenshotDir,
		VideoTrack: videoTrack,
	}
}

func (s *VideoServer) StreamFromDrone(drone *tello.Drone) {
	videoConn, err := drone.StreamVideo()

	if err != nil {
		fmt.Println("Error listening to video stream " + err.Error())
		return
	}

	h264Reader, _ := h264reader.NewReader(videoConn)

	for {
		nal, _ := h264Reader.NextNAL()

		s.Mu.Lock()
		s.nalBuffer = append(s.nalBuffer, nal.Data)
		if len(s.nalBuffer) > s.maxBufferSize {
				s.nalBuffer = s.nalBuffer[1:]
		}
		s.Mu.Unlock()

		s.VideoTrack.WriteSample(
			media.Sample{Data: nal.Data, Duration: time.Millisecond * 33},
		)
	}
}

func (s *VideoServer) SaveScreenshot() (string, error) {
	s.Mu.RLock()
	nals := make([][]byte, len(s.nalBuffer))
	copy(nals, s.nalBuffer)
	s.Mu.RUnlock()

	if len(nals) == 0 {
		return "", fmt.Errorf("No NALs in buffer, cannot take screenshot")
	}

	var h264Data bytes.Buffer
	for _, nal := range nals {
			h264Data.Write([]byte{0x00, 0x00, 0x00, 0x01})
			h264Data.Write(nal)
	}

	timestamp := time.Now().Unix()
	filename := filepath.Join(s.screenshotDir, fmt.Sprintf("screenshot_%d.jpg", timestamp))

	cmd := exec.Command("ffmpeg",
		"-loglevel", "error",
		"-sseof", "-1",
		"-f", "h264",
		"-i", "pipe:0",
		"-vframes", "1",
		"-q:v", "1",
		"-y",
		filename,
	)
    
  cmd.Stdin = &h264Data

	var stderr bytes.Buffer
	cmd.Stderr = &stderr
    
	if err := cmd.Run(); err != nil {
	  return "", fmt.Errorf("ffmpeg error: %v - %s", err, stderr.String())
	}

	return filename, nil
}
