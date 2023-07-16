package pkg

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type tui struct {
	app *tview.Application
}

func JobList() {
	app := tview.NewApplication()
	t := &tui{app: app}
	jobs := tview.NewList().
		ShowSecondaryText(false).
		SetBorder(true).
		SetTitle("채용 회사 / 직무")
	jobDescriptions := tview.NewTable().
		SetBorder(true).
		SetTitle("JD")

	flex := tview.NewFlex().
		AddItem(jobs, 0, 1, true).
		AddItem(jobDescriptions, 0, 2, false)

	app.SetInputCapture(t.inputActions)
	app.SetRoot(flex, true).Run()
}

func (t *tui) inputActions(e *tcell.EventKey) *tcell.EventKey {
	switch pressedKey := e.Rune(); pressedKey {
	case 'q':
		t.app.Stop()
	}
	return e
}
