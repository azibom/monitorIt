package path

import (
	"os"
)

func RootPath() string {
	currentPath, _ := os.Getwd()
	rootPath := currentPath + "/../../"
	return rootPath
}

func StaticPath() string {
	staticPath := RootPath() + "ui/static/"
	return staticPath
}

func UiPath() string {
	uiPath := RootPath() + "ui/"
	return uiPath
}
