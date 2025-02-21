package main

import "fmt"

func commandGetCharacters(cfg *config, args ...string) error {
	charResp, err := cfg.theoneapiClient.ListCharacters()
	if err != nil {
		return err
	}

	fmt.Println("\n---List of Characters---")
	for _, character := range charResp.Docs {
		fmt.Println(" -", character.Name)
	}
	fmt.Println("")
	return nil
}
