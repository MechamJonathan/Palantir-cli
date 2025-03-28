package styles

import (
	"github.com/charmbracelet/lipgloss"
)

var Title = lipgloss.NewStyle().
	Bold(true).
	Width(70).
	Align(lipgloss.Center).
	PaddingTop(2).
	PaddingRight(4).
	PaddingBottom(1).
	PaddingLeft(4).
	Foreground(Yellow)

var SubHeader = lipgloss.NewStyle().
	Align(lipgloss.Center).
	Width(70).
	PaddingBottom(1).
	PaddingTop(1).
	Align(lipgloss.Center).
	Faint(true)

var StartUpQuote = lipgloss.NewStyle().
	Align(lipgloss.Center).
	Width(70).
	Margin(10, 2, 10, 2).
	Align(lipgloss.Center).
	Faint(true).Bold(true)
