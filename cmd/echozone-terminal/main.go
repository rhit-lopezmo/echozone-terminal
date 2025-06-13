package main

import (
	"echozone-terminal/internal/controls"
	"echozone-terminal/internal/logging"
	"echozone-terminal/internal/ui"

	"github.com/rivo/tview"
)

func main() {
	logging.InitializeLogger()
	app := tview.NewApplication()
	layout := ui.CreateLayout()
	controls.SetupControls(app, layout)

	err := app.SetRoot(layout.Root, true).
						 EnableMouse(true).
						 Run()

	if err != nil {
		panic(err)
	}
}

