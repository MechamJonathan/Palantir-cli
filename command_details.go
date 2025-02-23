package main

import (
	"errors"
	"fmt"
	"strings"
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
	if movieErr == nil {
		return nil
	}

	bookErr := fetchBookDetails(cfg, inputName)
	if movieErr == nil {
		return nil
	}

	if movieErr != nil && bookErr != nil {
		return fmt.Errorf("no details found for IIII: %s", inputName)
	}

	return fmt.Errorf("no details found for: %s", inputName)
}

func fetchMovieDetails(cfg *config, name string) error {
	movieResp, err := cfg.theoneapiClient.GetMovieByName(name)
	if err != nil {
		return err
	}

	fmt.Printf("\n%-20s\n", "Movie Details")
	fmt.Println("--------------------")
	fmt.Printf(" - Name: %s\n - ID: %s\n - Runtime: %d min\n - Budget: $%.2fM\n - Box Office: $%.2fM\n - Awards: %d nominations, %d wins\n - Rotten Tomatoes: %.1f%%\n\n",
		movieResp.Name, movieResp.ID, movieResp.RuntimeInMinutes, movieResp.BudgetInMillions,
		movieResp.BoxOfficeRevenueInMillions, movieResp.AcademyAwardNominations, movieResp.AcademyAwardWins,
		movieResp.RottenTomatoesScore)

	return nil
}

func fetchBookDetails(cfg *config, name string) error {
	bookResp, err := cfg.theoneapiClient.GetBookByName(name)
	if err != nil {
		return err
	}

	fmt.Printf("\n%-20s\n", "Book Details")
	fmt.Println("--------------------")
	fmt.Printf(" - Name: %s\n - No additional details available currently\n\n", bookResp.Name)

	return nil
}

func fetchCharacterDetails(cfg *config, name string) error {
	charResp, err := cfg.theoneapiClient.GetCharacterByName(name)
	if err != nil {
		return err
	}

	fmt.Printf("\n%-20s\n", "Character Details")
	fmt.Println("--------------------")
	fmt.Printf(" - Name: %s\n - WikiURL: %s\n - Race: %s\n - Birth: %s\n - Gender: %s\n - Death: %s\n - Hair: %s\n - Height: %s\n - Realm: %s\n - Spouse: %s\n\n",
		charResp.Name, charResp.WikiURL, charResp.Race, charResp.Birth, charResp.Gender, charResp.Death, charResp.Hair, charResp.Height, charResp.Realm, charResp.Spouse)

	return nil
}
