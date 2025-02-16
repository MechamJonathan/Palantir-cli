package main

import "fmt"

func getMovies(cfg *config) error {
	movieResp, err := cfg.theoneapiClient.ListMovies()
	if err != nil {
		return err
	}

	fmt.Println("\n--List of Movies--")
	for _, movie := range movieResp.Docs {
		fmt.Println(" -", movie.Name, " id:", movie.ID)
	}
	fmt.Println("")
	return nil
}
