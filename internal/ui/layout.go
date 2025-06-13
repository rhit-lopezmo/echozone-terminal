package ui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type Layout struct {
	NowPlaying *tview.TextView
	InputBar *tview.InputField
	OutputLog *tview.TextView
	Root *tview.Flex
}

const FOCUS_BORDER_COLOR = tcell.ColorYellow
const BLUR_BORDER_COLOR = tcell.ColorGrey

func CreateLayout() *Layout {
	nowPlaying := tview.NewTextView().
											SetText("Now Playing: ...").
											SetDynamicColors(true)

	nowPlaying.SetBorder(true)	
	nowPlaying.SetBorderColor(BLUR_BORDER_COLOR)

	nowPlaying.SetFocusFunc(func() {
		nowPlaying.SetBorderColor(FOCUS_BORDER_COLOR)
	})

	nowPlaying.SetBlurFunc(func() {
		nowPlaying.SetBorderColor(BLUR_BORDER_COLOR)
	})
	
	inputBar := tview.NewInputField().
										SetText("").
										SetLabelWidth(0)
	
	inputBar.SetBorder(true)	
	inputBar.SetBorderColor(BLUR_BORDER_COLOR)

	outputLog := tview.NewTextView().
										 SetDynamicColors(true).
										 SetScrollable(true)

	outputLog.SetBorder(true) 
	outputLog.SetBorderColor(BLUR_BORDER_COLOR)

	outputLog.SetFocusFunc(func() {
		outputLog.SetBorderColor(FOCUS_BORDER_COLOR)
	})

	outputLog.SetBlurFunc(func() {
		outputLog.SetBorderColor(BLUR_BORDER_COLOR)
	})

	root := tview.NewFlex().
								SetDirection(tview.FlexRow).
								AddItem(nowPlaying, 0, 1, true).
								AddItem(inputBar, 3, 0, true).
								AddItem(outputLog, 0, 5, true)

	return &Layout {
		NowPlaying: nowPlaying,
		InputBar: inputBar,
		OutputLog: outputLog,
		Root: root,
	}
}
