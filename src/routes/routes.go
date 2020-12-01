package routes

import (
	"controllers/homeController"
	"helpers/path"
	"net/http"
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
	fs := http.FileServer(http.Dir(path.StaticPath()))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
}
