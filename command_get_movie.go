package main

import (
	"errors"
	"fmt"
	"strings"
)

func commandGetMovie(cfg *config, args ...string) error {
	movieName := strings.Join(args, " ")
	//fmt.Println("Args received:", args)

	if len(args) < 1 {
		return errors.New("you must provide a Movie name")
	}

	movieResp, err := cfg.theoneapiClient.GetMovieByName(movieName)
	if err != nil {
		return err
	}
	//fmt.Println("Movie response received:", movieResp)

	fmt.Printf("\n---Movie Details---\n")
	fmt.Printf("Name: %s\n", movieResp.Name)
	fmt.Printf("Runtime: %d minutes\n", movieResp.RuntimeInMinutes)
	fmt.Printf("Budget: $%.2f million\n", movieResp.BudgetInMillions)
	fmt.Printf("Box Office Revenue: $%.2f million\n", movieResp.BoxOfficeRevenueInMillions)
	fmt.Printf("Academy Award Nominations: %d\n", movieResp.AcademyAwardNominations)
	fmt.Printf("Academy Award Wins: %d\n", movieResp.AcademyAwardWins)
	fmt.Printf("Rotten Tomatoes Score: %.2f%%\n", movieResp.RottenTomatoesScore)
	fmt.Println("")
	return nil
}
