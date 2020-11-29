package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"text/template"
	"time"

	"github.com/gorilla/websocket"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func homePage(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("index.html")
	tmpl.Execute(w, "data goes here")
}

type SystemInfo struct {
	CPU    []float64
	MEMORY float64
}

func wsEndpoint(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}

	log.Println("Client Connected")
	if err != nil {
		log.Println(err)
	}

	for {
		percentage, _ := cpu.Percent(0, true)
		swapMemory, _ := mem.SwapMemory()
		systemInfo := SystemInfo{percentage, swapMemory.UsedPercent}
		jsonData, err := json.Marshal(systemInfo)
		if err != nil {
			fmt.Println(err)
			return
		}

		ws.WriteMessage(1, jsonData)
		time.Sleep(1000 * time.Millisecond)
	}
}

func setupRoutes() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/ws", wsEndpoint)

	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
}

func main() {
	fmt.Println("Hello World")
	setupRoutes()
	log.Fatal(http.ListenAndServe(":1234", nil))
}
