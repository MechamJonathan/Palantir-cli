package main

import (
	"fmt"
	"os"

	"github.com/MechamJonathan/lotr-companion-app/styles"
)

func commandExit(cfg *config, args ...string) error {
	if err := clearScreen(); err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	fmt.Println(styles.SubHeader.Render("Closing Palantír... Namárië!"))
	MoveCursorToBottom()
	os.Exit(0)
	return nil
}
