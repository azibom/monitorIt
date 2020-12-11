package main

import (
	"fmt"
	"log"
	"net/http"

	"monitoring/internal/package/routes"
)

func main() {
	routes.Init()
	fmt.Println("Listening at port 1234 ...")
	log.Fatal(http.ListenAndServe(":1234", nil))
}
