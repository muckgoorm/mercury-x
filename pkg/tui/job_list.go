package tui

import (
	"mercury-x/internal"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type jobList struct {
	list list.Model
}

type job struct {
	company string
	role    string
}

func (j job) Title() string       { return j.company }
func (j job) Description() string { return j.role }
func (j job) FilterValue() string { return j.company }

func (jl jobList) Init() tea.Cmd {
	return nil
}

func (jl jobList) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "esc", "ctrl+c":
			return jl, tea.Quit
		}
	case tea.WindowSizeMsg:
		h, v := listStyle.GetFrameSize()
		jl.list.SetSize(msg.Width-h, msg.Height-v)
	}

	var cmd tea.Cmd
	jl.list, cmd = jl.list.Update(msg)
	return jl, cmd
}

func (jl jobList) View() string {
	return listStyle.Render(jl.list.View())
}

func JobList(postings []internal.JobPosting) error {
	var items []list.Item
	for _, posting := range postings {
		items = append(items, job{company: posting.Company, role: posting.Role})
	}

	m := jobList{list: list.New(items, list.NewDefaultDelegate(), 0, 0)}
	m.list.Title = "채용 공고"

	p := tea.NewProgram(m, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		return err
	}

	return nil
}
