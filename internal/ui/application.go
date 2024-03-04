package ui

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"log"
)

func NewApplication(version, name string) error {
	a := app.New()
	w := a.NewWindow(name)

	log.Printf("App builded on version: %s", version)

	hello := widget.NewLabel("Hello Fyne !")
	w.SetContent(container.NewVBox(
		hello,
		widget.NewButton("Hi !", func() {
			hello.SetText("Welcome :)")
		}),
	))

	w.ShowAndRun()

	return nil
}
