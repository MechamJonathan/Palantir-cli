package main

import (
	"fmt"
	"os"

	"github.com/MechamJonathan/lotr-companion-app/styles"
)

func commandExit(cfg *config, args ...string) error {
	clearScreen()
	fmt.Println(styles.SubHeader.Render("Closing Lotr-Companion-App... Goodbye!"))
	MoveCursorToBottom()
	os.Exit(0)
	return nil
}
