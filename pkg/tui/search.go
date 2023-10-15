package tui

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var listStyle = lipgloss.NewStyle().Margin(1, 2)

type searchState int

const (
	stateJobList searchState = iota
	stateJobDetail
)

func (s searchState) String() string {
	return map[searchState]string{
		stateJobList:   "showing job list",
		stateJobDetail: "showing job description",
	}[s]
}

type searchModel struct {
	state    searchState
	jobList  jobList
	jobPager jobPager
}

func (s searchModel) Init() tea.Cmd {
	var cmds []tea.Cmd
	s.jobList = jobList{list: list.New(nil, list.NewDefaultDelegate(), 0, 0)}
	s.jobList.list.Title = "채용 공고"
	s.jobPager = jobPager{pager: viewport.Model{}}

	cmds = append(cmds, s.jobList.Init())

	return nil
}

func (s searchModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			if s.state == stateJobDetail {
				// TODO: unload job detail
			}
		case "q":
			if s.state == stateJobList {
				return s, tea.Quit
			}
		case "ctrl+c":
			return s, tea.Quit
		}
	case tea.WindowSizeMsg:
		s.jobList.list.SetSize(msg.Width, msg.Height)
		s.jobPager.pager.Width = msg.Width
		s.jobPager.pager.Height = msg.Height
	}

	return s, tea.Batch(cmds...)
}

func (s searchModel) View() string {
	return ""
}
