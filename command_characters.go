package main

import (
	"errors"
	"fmt"

	"github.com/MechamJonathan/lotr-companion-app/internal/theoneapi"
)

func commandGetCharacters(cfg *config, args ...string) error {
	if len(args) < 1 {
		return errors.New("usage: characters <all>|<fellowship>|<hobbits>")
	}

	charResp, err := cfg.theoneapiClient.ListCharacters()
	if err != nil {
		return err
	}

	if args[0] == "all" {
		fmt.Println("\n--- List of Characters ---")
		for _, character := range charResp.Docs {
			fmt.Println(" -", character.Name)
		}
	} else if args[0] == "fellowship" {
		getFellowshipMembers(charResp.Docs)
	} else {
		return fmt.Errorf("invalid argument '%s'", args[0])
	}

	fmt.Println("")
	return nil
}

func getFellowshipMembers(characters []theoneapi.Character) {
	var fellowshipMembers = []string{
		"Frodo Baggins", "Samwise Gamgee", "Gandalf", "Aragorn II Elessar", "Legolas",
		"Gimli", "Boromir", "Meriadoc Brandybuck", "Peregrin Took",
	}

	fmt.Println("\n--- Fellowship Members ---")
	for _, character := range characters {
		for _, member := range fellowshipMembers {
			if character.Name == member {
				fmt.Println(" -", character.Name)
				break
			}
		}
	}
}
