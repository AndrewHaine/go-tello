package main

import (
	"fmt"
	"net/http"
)

const (
  staticDir = "cmd/web-app/web/dist"
  port = "8080"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("./" + staticDir)))

	fmt.Println("Starting web server on port " + port + "; serving files from " + staticDir)
	err := http.ListenAndServe(":" + port, nil)

	if err != nil {
		fmt.Println("Error starting http server: " + err.Error())
	}
}
