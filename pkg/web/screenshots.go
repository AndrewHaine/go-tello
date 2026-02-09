package web

import (
	"encoding/json"
	"net/http"
	"os"
)

type Screenshot struct {
	Name string `json:"name"`
	Date string `json:"date"`
}

type ScreenshotRes struct {
	Screenshots []Screenshot `json:"screenshots"`
}

func TakeScreenshot(h *Hub) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {
		filename, err := h.VideoServer.SaveScreenshot()

		if err != nil {
			w.Write([]byte("Error taking screenshot " + err.Error()))
			return
		}

		w.Write([]byte("Success!"))

		// Broadcast a message to notify the browsers that we have a new screenshot
		h.Broadcast <- Event{
			Event: EventTypeScreenshotAdded,
			Payload: map[string]any {
				"file": filename,
			},
		}
	}
}

func ServeScreenshots(dir string) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {
		screenshots := []Screenshot{}
		entries, _ := os.ReadDir(dir)

		for _, entry := range entries {
			info, _ := entry.Info()

			if entry.Name() == ".gitignore" {
				continue
			}

			screenshots = append(screenshots, Screenshot{
				Name: "/screenshots/" + entry.Name(),
				Date: info.ModTime().String(),
			})
		}

		res := ScreenshotRes{
			Screenshots: screenshots,
		}
		marshalledRes, _ := json.Marshal(res)

		w.Write(marshalledRes)
		w.Header().Add("Content-Type", "application/json")
	}
}
