package main

import (
	"fmt"

	"github.com/MechamJonathan/lotr-companion-app/styles"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
)

func commandGetBooks(cfg *config, args ...string) error {
	booksResp, err := cfg.theoneapiClient.ListBooks()
	if err != nil {
		return err
	}

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
		Headers("Books").Width(72)
	for _, book := range booksResp.Docs {
		t.Row(book.Name)
	}
	fmt.Println("")
	fmt.Println(t)

	// fmt.Printf("\n%-20s\n", "Books")
	// fmt.Println("--------------------")
	// for _, book := range booksResp.Docs {
	// 	fmt.Println(" -", book.Name)
	// }
	fmt.Println("")
	return nil
}
