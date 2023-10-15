package tui

import (
	"mercury-x/internal"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
)

var cmd tea.Cmd

type jobPager struct {
	pager viewport.Model
	jd    internal.JobDescription
}

func (jp jobPager) Init() tea.Cmd {
	return nil
}

func (jp jobPager) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		top, right, bottom, left := docStyle.GetMargin()
		jp.pager = viewport.New(windowSize.Width-left-right, windowSize.Height-top-bottom)
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, keys.Back):
			// TODO: unload job detail and return to job list
		case key.Matches(msg, keys.Quit):
			return jp, tea.Quit
		}
	}

	jp.pager, cmd = jp.pager.Update(msg)
	// TODO: set job detail content
	return jp, cmd
}

func (jp jobPager) View() string {
	return docStyle.Render(jp.pager.View())
}
