package styles

import "github.com/charmbracelet/lipgloss"

var Header = lipgloss.NewStyle().
	Bold(true).
	Width(63).
	Align(lipgloss.Center)

var SubHeader = lipgloss.NewStyle().
	Align(lipgloss.Center).
	Width(63).
	Align(lipgloss.Center).
	Faint(true)

var TableHeader = lipgloss.NewStyle().
	Bold(true).
	Align(lipgloss.Center)

var TableText = lipgloss.NewStyle().
	Align(lipgloss.Left)

var myCuteBorder = lipgloss.Border{
	Top:         "._.:*:",
	Bottom:      "._.:*:",
	Left:        "|*",
	Right:       "|*",
	TopLeft:     "*",
	TopRight:    "*",
	BottomLeft:  "*",
	BottomRight: "*",
}

var BorderHeader = lipgloss.NewStyle().BorderStyle(lipgloss.RoundedBorder()).
	BorderForeground(lipgloss.Color("228")).
	BorderBackground(lipgloss.Color("63")).
	BorderTop(true).
	BorderLeft(true)
