package path

import (
	"log"
	"os"
)

func RootPath() string {
	rootPath, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	rootPath = rootPath + "/"
	return rootPath
}

func AssetsPath() string {
	assetsPath := RootPath() + "assets/"
	return assetsPath
}

func JsPath() string {
	jsPath := AssetsPath() + "js/"
	return jsPath
}

func TemplatesPath() string {
	templatesPath := AssetsPath() + "templates/"
	return templatesPath
}
