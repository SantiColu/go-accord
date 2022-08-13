package tui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type TUI struct {
	App      *tview.Application
	Channels *tview.List
	Users    *tview.List
	Chat     *tview.List
	Input    *tview.InputField
	Grid     *tview.Grid
}

func (tui *TUI) Start() error {
	return tui.App.Run()
}

func (tui *TUI) setupKeyboard() {
	focusMapping := map[tview.Primitive]struct{ next, prev tview.Primitive }{
		tui.Channels: {tui.Chat, tui.Users},
		tui.Chat:     {tui.Input, tui.Channels},
		tui.Input:    {tui.Users, tui.Chat},
		tui.Users:    {tui.Channels, tui.Users},
	}

	// Setup app level keyboard shortcuts.
	tui.App.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyCtrlA:
			tui.App.SetFocus(tui.Channels)
		case tcell.KeyCtrlG:
			tui.App.SetFocus(tui.Chat)
		case tcell.KeyCtrlK:
			tui.App.SetFocus(tui.Users)
		case tcell.KeyCtrlB:
			tui.App.SetFocus(tui.Input)
		// On Tab press set focus to the next element.
		case tcell.KeyTab:
			if focusMap, ok := focusMapping[tui.App.GetFocus()]; ok {
				tui.App.SetFocus(focusMap.next)
			} else {
				tui.App.SetFocus(tui.Input)
			}

			// Return `nil` to avoid default Backtab behaviour for the primitive.
			return nil
		}
		return event
	})

}

func NewTUI() *TUI {
	t := TUI{}
	t.App = tview.NewApplication()

	// Setup view elements.
	t.Channels = tview.NewList().ShowSecondaryText(false)
	t.Users = tview.NewList().ShowSecondaryText(false)
	t.Chat = tview.NewList().ShowSecondaryText(false)
	t.Input = tview.NewInputField()

	// Configure appearance.
	t.Channels.SetBorder(true).SetTitle(" Channels ")
	t.Chat.SetBorder(true).SetTitle(" Channel Name ")
	t.Users.SetBorder(true).SetTitle(" Users ")
	t.Input.SetBorder(true)
	t.Input.SetPlaceholder("Enviar mensaje a #Channel Name").SetFieldBackgroundColor(tcell.ColorPurple)

	chat := tview.NewGrid().
		SetRows(0, 3).
		AddItem(t.Chat, 0, 0, 1, 1, 0, 0, false).
		AddItem(t.Input, 1, 0, 1, 1, 0, 0, false)

	t.Grid = tview.NewGrid().
		SetRows(0, 0, 0).
		SetColumns(25, 0, 25).
		AddItem(t.Channels, 0, 0, 3, 1, 0, 0, false).
		AddItem(chat, 0, 1, 3, 1, 0, 0, false).
		AddItem(t.Users, 0, 2, 3, 1, 0, 0, false)

	t.Users.AddItem("Santi", "", 0, nil)
	t.Users.AddItem("Facu", "", 0, nil)
	t.Channels.AddItem("General", "", 0, nil)
	t.Channels.AddItem("Juegos", "", 0, nil)

	t.setupKeyboard()

	t.App.SetRoot(t.Grid, true).EnableMouse(true)

	return &t
}
