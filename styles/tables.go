package styles

import "github.com/charmbracelet/lipgloss"

var (
	HeaderStyle  = lipgloss.NewStyle().Foreground(Orange).Bold(true).Align(lipgloss.Center)
	CellStyle    = lipgloss.NewStyle().Padding(0, 1)
	OddRowStyle  = CellStyle.Foreground(LightGray)
	EvenRowStyle = CellStyle.Foreground(Gray)

	OddQuoteStyle = OddRowStyle.PaddingBottom(2)
	QuoteStyle    = EvenRowStyle.PaddingBottom(2)

	MyCuteBorder = lipgloss.Border{
		Top:         "._.:*:",
		Bottom:      "._.:*:",
		Left:        "|*",
		Right:       "|*",
		TopLeft:     "*",
		TopRight:    "*",
		BottomLeft:  "*",
		BottomRight: "*",
	}
)
