package main

import (
	"errors"
	"fmt"

	"github.com/MechamJonathan/lotr-companion-app/internal/theoneapi"
)

func commandGetCharacters(cfg *config, args ...string) error {
	if len(args) < 1 {
		return errors.New("usage: characters <all> | <fellowship> | <hobbits> | <men> | " +
			"<elves> | <dwarves> | <orcs> | <wizards> | <creatures>")
	}

	charResp, err := cfg.theoneapiClient.ListCharacters()
	if err != nil {
		return err
	}

	switch args[0] {
	case "all":
		printAllCharacters(charResp.Docs)
	case "fellowship":
		getFellowshipMembers(charResp.Docs)
	case "hobbits":
		getHobbitMembers(charResp.Docs)
	case "men":
		getMenOfMiddleEarth(charResp.Docs)
	case "elves":
		getElves(charResp.Docs)
	case "dwarves":
		getDwarves(charResp.Docs)
	case "orcs":
		getOrcs(charResp.Docs)
	case "wizards":
		getWizards(charResp.Docs)
	case "creatures":
		getCreatures(charResp.Docs)
	default:
		return fmt.Errorf("invalid argument '%s'", args[0])
	}

	fmt.Println("")
	return nil
}

func printAllCharacters(characters []theoneapi.Character) {
	printHeader("All Characters")
	for _, character := range characters {
		fmt.Println(" -", character.Name)
	}
}

func getFellowshipMembers(characters []theoneapi.Character) {
	fellowshipMembers := []string{
		"Frodo Baggins", "Samwise Gamgee", "Gandalf", "Aragorn II Elessar", "Legolas",
		"Gimli", "Boromir", "Meriadoc Brandybuck", "Peregrin Took",
	}
	printGroupMembers("Fellowship Members", fellowshipMembers, characters)
}

func getHobbitMembers(characters []theoneapi.Character) {
	hobbitMembers := []string{
		"Frodo Baggins", "Samwise Gamgee", "Meriadoc Brandybuck", "Peregrin Took", "Bilbo Baggins",
	}
	printGroupMembers("Hobbits", hobbitMembers, characters)
}

func getMenOfMiddleEarth(characters []theoneapi.Character) {
	menOfMiddleEarth := []string{
		"Aragorn II Elessar", "Boromir", "Faramir", "Théoden", "Éomer", "Éowyn", "Denethor", "Bard",
		"Gríma Wormtongue", "Denethor II",
	}
	printGroupMembers("Men of Middle Earth", menOfMiddleEarth, characters)
}

func getElves(characters []theoneapi.Character) {
	elves := []string{
		"Legolas", "Elrond", "Galadriel", "Arwen", "Thranduil", "Glorfindel", "Haldir", "Celeborn",
	}
	printGroupMembers("Elves", elves, characters)
}

func getDwarves(characters []theoneapi.Character) {
	dwarves := []string{
		"Gimli", "Thorin II Oakenshield", "Balin", "Dwalin", "Bofur", "Durin", "Dáin II Ironfoot",
		"Fíli and Kíli", "Óin", "Glóin", "Bifur", "Bombur", "Dori", "Nori", "Ori",
	}
	printGroupMembers("Dwarves", dwarves, characters)
}

func getWizards(characters []theoneapi.Character) {
	wizards := []string{
		"Gandalf", "Saruman", "Sauron", "Radagast", "Alatar", "Pallando",
	}
	printGroupMembers("Wizards", wizards, characters)
}

func getOrcs(characters []theoneapi.Character) {
	orcs := []string{
		"Azog", "Bolg", "Gothmog", "Uglúk", "Grishnákh", "Shagrat", "Gorbag", "Snaga",
	}
	printGroupMembers("Orcs", orcs, characters)
}

func getCreatures(characters []theoneapi.Character) {
	creatures := []string{
		"Gollum", "Smaug", "Shelob", "Treebeard", "Watcher in the Water", "Gwaihir", "Durin's Bane", "Witch-King of Angmar",
		"Khamûl",
	}
	printGroupMembers("Creatures", creatures, characters)
}

func printGroupMembers(title string, groupMembers []string, characters []theoneapi.Character) {
	printHeader(title)
	for _, character := range characters {
		for _, member := range groupMembers {
			if character.Name == member {
				fmt.Println(" -", character.Name)
				break
			}
		}
	}
}

// func printHeader(title string) {
// 	fmt.Printf("\n%-20s\n", title)
// 	fmt.Println("--------------------")
// }
