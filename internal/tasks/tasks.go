package tasks

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/d4nz-jpg/taskbrancher/style"
	"strings"
)

type status int

const (
	Todo status = iota
	InProgress
	Done
)

type Task struct {
	ID          int
	Title       string
	Description string
	Status      status
	Due         int64
	Subtasks    []Task
}

func (t Task) View() string {
	header := style.Header.Render(t.Title)
	description := style.Text.Render(t.Description)

	var list []string
	for _, element := range t.Subtasks {
		bullet := element.Bullet()
		list = append(list, bullet)
	}

	return lipgloss.JoinVertical(0,
		header,
		description,
		style.Separator.Render(strings.Join(list, "\n")),
	)
}

func (t Task) Bullet() string {
	return lipgloss.JoinHorizontal(0, "["+t.StatusIcon()+"] ", t.Title)
}

func (t Task) StatusIcon() string {
	switch t.Status {
	case Todo:
		return "✗"
	case Done:
		return "✓"
	case InProgress:
		return "◐"
	default:
		return " "
	}
}
