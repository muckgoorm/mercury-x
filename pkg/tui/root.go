package tui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type tui struct {
	app *tview.Application
}

func newTui() *tui {
	return &tui{
		app: tview.NewApplication(),
	}
}

func (t *tui) run() {
	t.app.Run()
}

func (t *tui) quit(e *tcell.EventKey) *tcell.EventKey {
	switch pressedKey := e.Rune(); pressedKey {
	case 'q':
		t.app.Stop()
	}
	return e
}
