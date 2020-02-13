package main

import (
	"fmt"
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
)

const title = "WeatherStats"

func createDataPointSelector() (*widget.Select, *widget.Box) {
	selector := widget.NewSelect([]string{
		"maxtemp",
		"mintemp",
		"meantemp",
		"rain",
		"snow",
		"snowgrnd",
		"precip",
		"maxgust",
	}, func(dataPoint string) {})

	selector.SetSelected("maxtemp")
	return selector, widget.NewHBox(
			widget.NewLabel("Select a Data Point:"),
			selector,
		)
}

func main() {
	a := app.New()

	dataPointSelector, dataPointBox := createDataPointSelector()

	w := a.NewWindow(title)
	w.SetContent(widget.NewVBox(
		widget.NewLabel(title),
		dataPointBox,
		widget.NewButton("Quit", func() {
			fmt.Println(dataPointSelector.Selected)
			a.Quit()
		}),
	))

	w.SetPadded(true)
	w.SetFullScreen(true)
	w.ShowAndRun()
}
