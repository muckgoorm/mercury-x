package tui

import "github.com/rivo/tview"

func GenerateFlexTable(firstTitle, secondTitle string, firstItems []string) {
	t := newTui()

	firstTab := tview.NewList().ShowSecondaryText(false)
	for _, item := range firstItems {
		firstTab.AddItem(item, "", 0, nil)
	}
	firstTab.SetBorder(true).SetTitle(firstTitle)

	secondTab := tview.NewTable().
		SetBorder(true).
		SetTitle(secondTitle)

	flex := tview.NewFlex().
		AddItem(firstTab, 0, 1, true).
		AddItem(secondTab, 0, 2, false)

	t.app.SetInputCapture(t.quit)
	t.app.SetRoot(flex, true)
	t.run()
}
