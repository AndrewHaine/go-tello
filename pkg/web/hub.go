package web

import (
	"fmt"
	"time"

	"github.com/andrewhaine/go-tello/pkg/tello"
	"github.com/pion/webrtc/v3"
	"github.com/pion/webrtc/v3/pkg/media"
	"github.com/pion/webrtc/v3/pkg/media/h264reader"
)

type Hub struct {
	Browsers map[Browser]bool

	Broadcast chan Event

	Commands chan []byte

	Register chan *Browser

	Deregister chan *Browser

	VideoTrack webrtc.TrackLocal
}

func NewHub() Hub {
	return Hub {
		Browsers: make(map[Browser]bool),
		Broadcast: make(chan Event),
		Commands: make(chan []byte),
		Register: make(chan *Browser),
		Deregister: make(chan *Browser),
	}
}

func (h *Hub) RemoveBrowser(browser *Browser) {
	delete(h.Browsers, *browser)
	close(browser.Queue)
}

func (h *Hub) Listen() {
	for {
		select {
		case browser := <-h.Register:
			h.Browsers[*browser] = true

		case browser := <-h.Deregister:
			if _, ok := h.Browsers[*browser]; ok {
				h.RemoveBrowser(browser)
			}

		case msg := <-h.Broadcast:
			for browser := range h.Browsers {
				select {
				case browser.Queue <- msg:
				default:
					h.RemoveBrowser(&browser)
				}
			}
		}
	}
}

func (h *Hub) ListenVideo(drone *tello.Drone) {
	videoTrack, _ := webrtc.NewTrackLocalStaticSample(
		webrtc.RTPCodecCapability{MimeType: webrtc.MimeTypeH264},
		"video",
		"stream",
	)

	h.VideoTrack = videoTrack

	videoConn, err := drone.StreamVideo()

	if err != nil {
		fmt.Println("Error listening to video stream " + err.Error())
		return
	}

	h264Reader, _ := h264reader.NewReader(videoConn)
	for {
		nal, _ := h264Reader.NextNAL()
		videoTrack.WriteSample(
			media.Sample{Data: nal.Data, Duration: time.Millisecond * 33},
		)
	}
}
