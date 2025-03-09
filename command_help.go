package main

import (
	"fmt"
	"sort"

	"github.com/MechamJonathan/lotr-companion-app/styles"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
)

func commandHelp(cfg *config, args ...string) error {
	fmt.Println()
	fmt.Println(styles.Header.Render("Welcome to Palantír!"))
	fmt.Println(styles.SubHeader.Render("The unofficial LOTR companion app"))

	commands := getCommands()

	var commandNames []string
	for name := range commands {
		commandNames = append(commandNames, name)
	}

	sort.Strings(commandNames)

	t := table.New().
		Border(lipgloss.RoundedBorder()).
		BorderStyle(lipgloss.NewStyle().Foreground(lipgloss.Color("238"))).
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
		Headers("Command", "Description")

	for _, name := range commandNames {
		cmd := commands[name]
		t.Row(cmd.name, cmd.description)
	}
	fmt.Println(t)

	// fmt.Printf("%-20s %s\n", "Command", "Description")
	// fmt.Println("-------------------- ------------------------------------------")

	// for _, name := range commandNames {
	// 	cmd := commands[name]
	// 	fmt.Printf("%-20s %s\n", cmd.name, cmd.description)
	// }

	fmt.Println()
	return nil
}
