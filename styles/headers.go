package styles

import (
	"github.com/charmbracelet/lipgloss"
)

var Header = lipgloss.NewStyle().
	Bold(true).
	Width(70).
	Align(lipgloss.Center).
	PaddingTop(2).
	PaddingRight(4).
	PaddingBottom(2).
	PaddingLeft(4)

var SubHeader = lipgloss.NewStyle().
	Align(lipgloss.Center).
	Width(70).
	PaddingBottom(1).
	Align(lipgloss.Center).
	Faint(true)

var TableHeader = lipgloss.NewStyle().
	Bold(true).
	Align(lipgloss.Center)

var TableText1 = lipgloss.NewStyle().
	Align(lipgloss.Left).
	PaddingLeft(1).
	PaddingBottom(1).
	BorderStyle(lipgloss.HiddenBorder()).
	BorderBackground(lipgloss.Color("#FF5E5B"))
