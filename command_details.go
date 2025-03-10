package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/MechamJonathan/lotr-companion-app/styles"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
)

func commandGetDetails(cfg *config, args ...string) error {
	if len(args) < 1 {
		return errors.New("usage: details <character name>|<movie>|<book>")
	}

	inputName := strings.Join(args, " ")

	err := fetchCharacterDetails(cfg, inputName)
	if err == nil {
		return nil
	}

	movieErr := fetchMovieDetails(cfg, inputName)
	bookErr := fetchBookDetails(cfg, inputName)
	if movieErr == nil || bookErr == nil {
		return nil
	}
	return fmt.Errorf("no details found for: %s", inputName)

}

func fetchMovieDetails(cfg *config, name string) error {
	movieResp, err := cfg.theoneapiClient.GetMovieByName(name)
	if err != nil {
		return err
	}

	// printHeader("Movie Details")
	// fmt.Printf(" - Name: %s\n - ID: %s\n - Runtime: %d min\n - Budget: $%.2fM\n - Box Office: $%.2fM\n - Awards: %d nominations, %d wins\n - Rotten Tomatoes: %.1f%%\n\n",
	// 	movieResp.Name, movieResp.ID, movieResp.RuntimeInMinutes, movieResp.BudgetInMillions,
	// 	movieResp.BoxOfficeRevenueInMillions, movieResp.AcademyAwardNominations, movieResp.AcademyAwardWins,
	// 	movieResp.RottenTomatoesScore)

	runtime := fmt.Sprint(movieResp.RuntimeInMinutes)
	budget := strconv.FormatFloat(movieResp.BudgetInMillions, 'f', -1, 64)
	boxOffice := strconv.FormatFloat(movieResp.BoxOfficeRevenueInMillions, 'f', -1, 64)
	rottenTomatoesScore := strconv.FormatFloat(movieResp.RottenTomatoesScore, 'f', -1, 64)
	awards := fmt.Sprintf("Awards: %d nominations, %d wins", movieResp.AcademyAwardNominations, movieResp.AcademyAwardWins)

	rows := [][]string{
		{"Name", movieResp.Name},
		{"Runtime", runtime + " mins"},
		{"Budget", "$" + budget + "M"},
		{"Box Office", "$" + boxOffice + "M"},
		{"Awards", awards},
		{"Rotten Tomatos", rottenTomatoesScore},
	}

	movieTableName := movieResp.Name + " (movie)"
	printDetailsTable(rows, movieTableName)
	return nil
}

func fetchBookDetails(cfg *config, name string) error {
	bookResp, err := cfg.theoneapiClient.GetBookByName(name)
	if err != nil {
		return err
	}

	// printHeader("Book Details")
	// fmt.Printf(" - Name: %s\n - No additional details available currently\n\n", bookResp.Name)

	rows := [][]string{
		{"Name", bookResp.Name},
		{"", "(No additional details availble currently)"},
	}

	bookTableName := bookResp.Name + " (book)"
	printDetailsTable(rows, bookTableName)
	return nil
}

func fetchCharacterDetails(cfg *config, name string) error {
	charResp, err := cfg.theoneapiClient.GetCharacterByName(name)
	if err != nil {
		return err
	}

	// printHeader("Character Details")
	// fmt.Printf(" - Name: %s\n - ID: %s\n - WikiURL: %s\n - Race: %s\n - Birth: %s\n - Gender: %s\n - Death: %s\n - Hair: %s\n - Height: %s\n - Realm: %s\n - Spouse: %s\n\n",
	// 	charResp.Name, charResp.ID, charResp.WikiURL, charResp.Race, charResp.Birth, charResp.Gender, charResp.Death, charResp.Hair, charResp.Height, charResp.Realm, charResp.Spouse)

	rows := [][]string{
		{"Name", charResp.Name},
		{"WikiURL", charResp.WikiURL},
		{"Race", charResp.Race},
		{"Birth", charResp.Birth},
		{"Gender", charResp.Gender},
		{"Death", charResp.Death},
		{"Hair", charResp.Hair},
		{"Height", charResp.Height},
		{"Realm", charResp.Realm},
		{"Spouse", charResp.Spouse},
	}

	printDetailsTable(rows, charResp.Name)
	return nil
}

func printDetailsTable(rows [][]string, name string) {

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
		Headers("", name).
		Width(72).
		Rows(rows...)

	fmt.Println("")
	fmt.Println(t)
	fmt.Println("")
}
