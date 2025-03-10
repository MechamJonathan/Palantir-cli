package main

import (
	"fmt"
	"os"
	"os/exec"

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
