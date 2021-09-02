package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("HHO GoShipify Dryrun Tool")
	w.SetFixedSize(true)
	w.Resize(
		fyne.Size{
			Width:  600,
			Height: 500,
		})

	hello := widget.NewTextGrid()
	hello.SetText("\n\n\n\n\n\n\n\n\n\n\n\nHello HHOer\n\n")
	w.SetContent(container.NewVBox(
		hello,
		widget.NewButton("Dryrun", func() {
			hello.SetText("Welcome :)")
		}),
	))

	w.ShowAndRun()
}
