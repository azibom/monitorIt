package websocketService

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"

	"monitoring/internal/package/services/systemInfoService"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func MonitoringReportsGenerator(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}

	log.Println("Client Connected")

	for {
		systemInfo := systemInfoService.GetSystemResourceInfo()

		jsonData, err := json.Marshal(systemInfo)
		if err != nil {
			fmt.Println(err)
			return
		}

		ws.WriteMessage(1, jsonData)
		time.Sleep(1000 * time.Millisecond)
	}
}
