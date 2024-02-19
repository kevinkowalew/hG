package list

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type SingleSelectList struct {
	height, selected int
	selectedStyle    lipgloss.Style
	items            []string
	quit             bool
}

func (l *SingleSelectList) Init() tea.Cmd {
	return nil
}

func (l *SingleSelectList) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "j":
			if l.selected < len(l.items)-1 {
				l.selected++
			}
		case "k":
			if l.selected > 0 {
				l.selected--
			}
		case "q", "ctrl+c":
			l.quit = true
			return l, tea.Quit
		case "enter":
			l.quit = true
			return l, tea.Quit
		}
	}

	return l, nil
}

func NewSingleSelectList(items []string, height int) *SingleSelectList {
	return &SingleSelectList{
		items:    items,
		selected: len(items) - 1,
		height:   height,
	}
}

func (l *SingleSelectList) View() string {
	if l.quit {
		return ""
	}

	start := max(0, len(l.items)-l.height)
	end := len(l.items)
	if l.selected < start {
		start = l.selected
		end = l.selected + l.height
	}

	view := ""
	for i := start; i < end; i++ {
		item := l.items[i]

		if i == l.selected {
			view += lipgloss.NewStyle().Foreground(lipgloss.Color("#874BFD")).Render(item)
		} else {
			view += item
		}

		if i < len(l.items)-1 {
			view += "\n"
		}
	}

	return view
}

func (l *SingleSelectList) Run() (string, error) {
	if _, err := tea.NewProgram(l).Run(); err != nil {
		return "", err
	}

	return l.items[l.selected], nil
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
