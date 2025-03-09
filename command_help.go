package main

import (
	"fmt"
	"sort"
)

func commandHelp(cfg *config, args ...string) error {
	fmt.Println()
	fmt.Println("Welcome to Palantír! The unofficial LOTR companion app.")
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println()

	commands := getCommands()

	var commandNames []string
	for name := range commands {
		commandNames = append(commandNames, name)
	}

	sort.Strings(commandNames)

	fmt.Printf("%-20s %s\n", "Command", "Description")
	fmt.Println("-------------------- ------------------------------------------")

	for _, name := range commandNames {
		cmd := commands[name]
		fmt.Printf("%-20s %s\n", cmd.name, cmd.description)
	}
	fmt.Println()
	return nil
}
