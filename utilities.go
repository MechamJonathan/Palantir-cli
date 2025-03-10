package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/MechamJonathan/lotr-companion-app/styles"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
	"golang.org/x/term"
)

func clearScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func MoveCursorToBottom() {
	lines := GetTerminalHeight()
	fmt.Printf("\033[%d;1H", lines)
}

func GetTerminalHeight() int {
	height, _, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		return 20
	}
	return height
}

func PrintUsageTable(cmdUsage string, options [][]string) {
	clearScreen()
	fmt.Println(styles.Title.Render(cmdUsage))
	t := table.New().
		Border(lipgloss.RoundedBorder()).
		BorderStyle(lipgloss.NewStyle().Foreground(lipgloss.Color(styles.Red))).
		StyleFunc(func(row, col int) lipgloss.Style {
			switch {
			case row == table.HeaderRow:
				return styles.HeaderStyle
			case row%2 == 0:
				return styles.EvenRowStyle
			default:
				return styles.OddRowStyle
			}
		}).
		Headers("Options", "Description").
		Width(72)

	t.Rows(options...)
	fmt.Println(t)
	MoveCursorToBottom()
}
