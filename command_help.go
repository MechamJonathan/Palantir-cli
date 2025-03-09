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
	fmt.Println()
	fmt.Println(styles.SubHeader.Render("The unofficial LOTR companion app."))

	commands := getCommands()

	var commandNames []string
	for name := range commands {
		commandNames = append(commandNames, name)
	}

	sort.Strings(commandNames)

	// fmt.Printf("%-20s %s\n", "Command", "Description")
	// fmt.Println("-------------------- ------------------------------------------")

	// for _, name := range commandNames {
	// 	cmd := commands[name]
	// 	fmt.Printf("%-20s %s\n", cmd.name, cmd.description)
	// }

	t := table.New().
		Border(lipgloss.NormalBorder()).
		BorderStyle(lipgloss.NewStyle().Foreground(lipgloss.Color("#FF5E5B"))).
		StyleFunc(func(row, col int) lipgloss.Style {
			switch {
			case row == table.HeaderRow:
				return styles.TableHeader
			case row%2 == 0:
				return styles.TableText
			default:
				return styles.TableText
			}
		}).
		Headers("Command", "Description")

	for _, name := range commandNames {
		cmd := commands[name]
		t.Row(cmd.name, cmd.description)
	}
	fmt.Println(t)

	fmt.Println()
	return nil
}
