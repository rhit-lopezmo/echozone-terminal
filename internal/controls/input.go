package controls

import (
	"echozone-terminal/internal/logging"
	"echozone-terminal/internal/ui"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func SetupControls(app *tview.Application, layout *ui.Layout) {
	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyRune {
			switch event.Rune() {
				case ':':
					logging.Debug(": was pressed, switching to input bar")
					app.SetFocus(layout.InputBar)
					break
				default:
					// logging.Error("unknown key pressed")
			}
		} else {
			switch event.Key() {
				case tcell.KeyEsc:
					logging.Debug("esc pressed, blurring current focus")	
					app.GetFocus().Blur()
					layout.InputBar.SetText("") // clear the input
					break
			}
		}

		return event
	})
}
