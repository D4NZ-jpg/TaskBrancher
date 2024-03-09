package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/d4nz-jpg/taskbrancher/internal/keys"
	"github.com/d4nz-jpg/taskbrancher/internal/tasks"
	"github.com/d4nz-jpg/taskbrancher/ui"
)

func main() {
	m := &ui.Model{
		Keys: keys.Keys,
		Task: tasks.Task{Status: tasks.Todo, Title: "Today", Description: "Today's tasks"},
	}
	p := tea.NewProgram(m, tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
