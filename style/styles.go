package style

import "github.com/charmbracelet/lipgloss"

var w = 48
var alt = "#2b2d42"
var accent = "#cdb4db"

var Separator = lipgloss.NewStyle().
	BorderStyle(lipgloss.RoundedBorder()).
	BorderForeground(lipgloss.Color(alt)).
	BorderTop(true).
	BorderBottom(true).
	Width(w)

var Title = lipgloss.NewStyle().
	Bold(true).
	PaddingTop(1).PaddingLeft(4).
	Width(w)

var Header = lipgloss.NewStyle().
	Width(w).
	Background(lipgloss.Color(accent)).
	Foreground(lipgloss.Color(alt)).
	Align(lipgloss.Center)

var Text = lipgloss.NewStyle().
	Width(w)

var Tip = lipgloss.NewStyle().
	Foreground(lipgloss.Color(alt))
