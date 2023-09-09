package tui

import (
	"mercury-x/internal"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var postingStyle = lipgloss.NewStyle().Margin(1, 2)

type posting struct {
	list list.Model
}

type job struct {
	company string
	role    string
}

func (j job) Title() string       { return j.company }
func (j job) Description() string { return j.role }
func (j job) FilterValue() string { return j.company }

func (p posting) Init() tea.Cmd {
	return nil
}

func (p posting) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "esc", "ctrl+c":
			return p, tea.Quit
		}
	case tea.WindowSizeMsg:
		h, v := postingStyle.GetFrameSize()
		p.list.SetSize(msg.Width-h, msg.Height-v)
	}

	var cmd tea.Cmd
	p.list, cmd = p.list.Update(msg)
	return p, cmd
}

func (p posting) View() string {
	return postingStyle.Render(p.list.View())
}

func JobList(postings []internal.JobPosting) error {
	var items []list.Item
	for _, posting := range postings {
		items = append(items, job{company: posting.Company, role: posting.Role})
	}

	m := posting{list: list.New(items, list.NewDefaultDelegate(), 0, 0)}
	m.list.Title = "채용 공고"

	p := tea.NewProgram(m, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		return err
	}

	return nil
}
