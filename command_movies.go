package main

import "fmt"

func commandGetMovies(cfg *config, args ...string) error {
	movieResp, err := cfg.theoneapiClient.ListMovies()
	if err != nil {
		return err
	}

	fmt.Printf("\n%-20s\n", "Movies")
	fmt.Println("--------------------")
	for _, movie := range movieResp.Docs {
		fmt.Println(" -", movie.Name)
	}
	fmt.Println("")
	return nil
}
