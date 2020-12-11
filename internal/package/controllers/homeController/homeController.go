package homeController

import (
	"html/template"
	"log"
	"net/http"

	"monitoring/internal/package/helpers/path"
	"monitoring/internal/package/services/websocketService"
)

func IndexAction(w http.ResponseWriter, r *http.Request) {

	tmpl, err := template.ParseFiles(path.TemplatesPath() + "index.html")
	if err != nil {
		log.Println(err)
	}
	tmpl.Execute(w, nil)
}

func WebSocketAction(w http.ResponseWriter, r *http.Request) {
	websocketService.MonitoringReportsGenerator(w, r)
}
