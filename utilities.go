package main

import (
	"fmt"
	"os"
	"os/exec"

	"golang.org/x/term"
)

func clearScreen() {
	cmd := exec.Command("clear") // Use "cls" on Windows
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func MoveCursorToBottom() {
	lines := GetTerminalHeight()
	fmt.Printf("\033[%d;1H", lines) // Move cursor to 'lines' row, column 1
}

func GetTerminalHeight() int {
	height, _, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		return 20 // Default fallback if unable to get size
	}
	return height
}
