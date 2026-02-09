package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/andrewhaine/go-tello/pkg/tello"
	"github.com/andrewhaine/go-tello/pkg/web"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

const (
  staticDirRel = "./web/dist"
	screenshotDirRel = "./screenshots"
  port = "8080"
)

var upgrader = websocket.Upgrader{}

func main() {
	drone := tello.NewDrone()

	err := drone.ConnectDefault()

  if err != nil {
    log.Printf("Could not connect to drone: %s", err)
    os.Exit(1)
  }

  defer drone.CloseConnection()
	
	drone.SendRawCmdString("command")
  
  hub := web.NewHub()
	go hub.Listen()

	webDir, _ := filepath.Abs(staticDirRel)
	screenshotDir, _ := filepath.Abs(screenshotDirRel)

	videoServer := web.NewVideoServer(screenshotDir)
	hub.VideoServer = &videoServer
	go videoServer.StreamFromDrone(&drone)

	go sendHubCommandsToDrone(&drone, &hub)
	go broadcastTelemetry(&drone, &hub)
	go broadcastMessages(&drone, &hub)

	r := mux.NewRouter()
	api := r.PathPrefix("/api").Subrouter()
	api.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(&hub, w, r)
	})
	api.HandleFunc("/screenshot", web.TakeScreenshot(&hub)).Methods("POST")
	api.HandleFunc("/screenshots", web.ServeScreenshots(screenshotDir)).Methods("GET")
	
	r.PathPrefix("/screenshots").Handler(http.StripPrefix("/screenshots/", http.FileServer(http.Dir(screenshotDir))))
	r.PathPrefix("/").Handler(http.FileServer(http.Dir(webDir)))
	
	srv := &http.Server{
		Handler: r,
		Addr: ":" + port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}

func serveWs(hub *web.Hub, w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}

	ws, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Println("Error upgrading connection", err.Error())
		return
	}

	log.Println("New browser connected!")

	browser := &web.Browser{Hub: hub, Conn: ws, Queue: make(chan web.Event, 256)}
	hub.Register <- browser

	go browser.SendQueuedMessages()
	go browser.ReceiveMessages()
}

func sendHubCommandsToDrone(drone *tello.Drone, hub *web.Hub) {
	for command := range hub.Commands {
		drone.SendRawCmdString(string(command))
	}
}

func broadcastTelemetry(drone *tello.Drone, hub *web.Hub) {
	telemetryChan, err := drone.StreamTelemetry()

	if err != nil {
		log.Println("Could not stream telemetry: " + err.Error())
		return
	}

	for telemetry := range telemetryChan {
		hub.Broadcast <- web.EventFromTelemetry(telemetry)
	}
}

func broadcastMessages(drone *tello.Drone, hub *web.Hub) {
	messageChan, err := drone.StreamMessages()

	if err != nil {
		log.Println("Could not stream messages: " + err.Error())
		return
	}

	for message := range messageChan {
		hub.Broadcast <- web.EventFromTelloMsg(message)
	}
}
