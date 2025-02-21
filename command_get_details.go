package main

import (
	"errors"
	"fmt"
	"strings"
)

func commandGetDetails(cfg *config, args ...string) error {
	if len(args) < 2 {
		return errors.New("usage: details [book|movie|character] <name>")
	}

	inputType := strings.ToLower(args[0])
	inputName := strings.Join(args[1:], " ")

	switch inputType {
	case "movie":
		return fetchMovieDetails(cfg, inputName)
	case "book":
		return fetchBookDetails(cfg, inputName)
	case "character":
		return fetchCharacterDetails(cfg, inputName)
	default:
		return errors.New("invalid type. Use 'details book <name>', 'details movie <name>', or 'details character <name>")
	}

}

func fetchMovieDetails(cfg *config, name string) error {
	movieResp, err := cfg.theoneapiClient.GetMovieByName(name)
	if err != nil {
		return err
	}

	fmt.Println("\n--- Movie Details ---")
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

	fmt.Println("\n--- Book Details ---")
	fmt.Printf(" - Name: %s\n - No additional details available currently\n\n", bookResp.Name)

	return nil
}

func fetchCharacterDetails(cfg *config, name string) error {
	charResp, err := cfg.theoneapiClient.GetCharacterByName(name)
	if err != nil {
		return err
	}

	fmt.Println("\n--- Character Details ---")
	fmt.Printf(" - Name: %s\n - WikiURL: %s\n - Race: %s\n - Birth: %s\n - Gender: %s\n - Death: %s\n - Hair: %s\n - Height: %s\n - Realm: %s\n - Spouse: %s\n\n",
		charResp.Name, charResp.WikiURL, charResp.Race, charResp.Birth, charResp.Gender, charResp.Death, charResp.Hair, charResp.Height, charResp.Realm, charResp.Spouse)

	return nil
}
