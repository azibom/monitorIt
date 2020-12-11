package routes

import (
	"net/http"

	"monitoring/internal/package/controllers/homeController"
	"monitoring/internal/package/helpers/path"
)

func Init() {
	handleRoutes()
	handleStaticFiles()
}

func handleRoutes() {
	http.HandleFunc("/", homeController.IndexAction)
	http.HandleFunc("/ws", homeController.WebSocketAction)
}

func handleStaticFiles() {
	fs := http.FileServer(http.Dir(path.AssetsPath()))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))
}
