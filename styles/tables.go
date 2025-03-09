package styles

import "github.com/charmbracelet/lipgloss"

var (
	HeaderStyle  = lipgloss.NewStyle().Foreground(Red).Bold(true).Align(lipgloss.Center)
	CellStyle    = lipgloss.NewStyle().Padding(0, 1)
	OddRowStyle  = CellStyle.Foreground(Gray)
	EvenRowStyle = CellStyle.Foreground(LightGray)
)
