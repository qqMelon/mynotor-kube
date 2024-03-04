package main

import (
	"github.com/qqMelon/mynotor-kube/internal/ui"
	"log"
	"os"
)

var (
	appName = "Mynotor Kube"
	version = "dev"
	commit  = "none"
	date    = "unknown"
	theme   = "dark"
)

func main() {
	os.Setenv("FYNE_THEME", theme)

	err := ui.NewApplication(version, appName)
	if err != nil {
		log.Fatal(err)
	}
}
