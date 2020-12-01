package homeController

import (
	"helpers/path"
	"html/template"
	"log"
	"net/http"
	"services/websocketService"
)

func IndexAction(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(path.UiPath() + "index.html")
	if err != nil {
		log.Println(err)
	}
	tmpl.Execute(w, nil)
}

func WebSocketAction(w http.ResponseWriter, r *http.Request) {
	websocketService.MonitoringReportsGenerator(w, r)
}
