package ui

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/d4nz-jpg/taskbrancher/internal/keys"
	"github.com/d4nz-jpg/taskbrancher/internal/tasks"
	"github.com/d4nz-jpg/taskbrancher/style"
)

type Model struct {
	index int
	Keys  keys.KeyMap
	Help  help.Model
	Task  tasks.Task
	Err   error
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.Help.Width = msg.Width

	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.Keys.Help):
			m.Help.ShowAll = !m.Help.ShowAll
		case key.Matches(msg, m.Keys.Quit):
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m Model) View() string {
	return lipgloss.JoinVertical(0, m.Task.View(), style.Tip.Render(m.Help.View(m.Keys)))
}
