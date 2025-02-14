package main

import (
	"fmt"
	"os"
)

func commandExit() error {
	fmt.Println("Closing Lotr-Companion-App... Goodbye!")
	os.Exit(0)
	return nil
}
