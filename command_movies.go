package main

import "fmt"

func commandGetMovies(cfg *config, args ...string) error {
	movieResp, err := cfg.theoneapiClient.ListMovies()
	if err != nil {
		return err
	}

	fmt.Println("\n---List of Movies---")
	for _, movie := range movieResp.Docs {
		fmt.Println(" -", movie.Name)
	}
	fmt.Println("")
	return nil
}
