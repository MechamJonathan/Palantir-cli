package main

import (
	"errors"
	"fmt"
	"strings"
)

func commandGetDetails(cfg *config, args ...string) error {
	if len(args) < 2 {
		return errors.New("usage: details [book|movie] <name>")
	}

	inputType := strings.ToLower(args[0])    // First argument specifies book or movie
	inputName := strings.Join(args[1:], " ") // Join remaining args as the name

	switch inputType {
	case "movie":
		movieResp, err := cfg.theoneapiClient.GetMovieByName(inputName)
		if err != nil {
			return err
		}
		fmt.Println("\n--- Movie Details ---")
		fmt.Printf("Name: %s\nID: %s\nRuntime: %d min\nBudget: $%.2fM\nBox Office: $%.2fM\nAwards: %d nominations, %d wins\nRotten Tomatoes: %.1f%%\n\n",
			movieResp.Name, movieResp.ID, movieResp.RuntimeInMinutes, movieResp.BudgetInMillions,
			movieResp.BoxOfficeRevenueInMillions, movieResp.AcademyAwardNominations, movieResp.AcademyAwardWins,
			movieResp.RottenTomatoesScore)

	case "book":
		bookResp, err := cfg.theoneapiClient.GetBookByName(inputName)
		if err != nil {
			return err
		}
		fmt.Println("\n--- Book Details ---")
		fmt.Printf("Name: %s\n\n", bookResp.Name)

	default:
		return errors.New("invalid type. Use 'details book <name>' or 'details movie <name>'")
	}

	return nil
}
