package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/MechamJonathan/lotr-companion-app/internal/theoneapi"
)

type config struct {
	theoneapiClient  theoneapi.Client
	currentQuotePage int
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

func startRepl(cfg *config) {
	reader := bufio.NewScanner(os.Stdin)
	cfg.currentQuotePage = 0

	for {
		fmt.Print("Lotr-Companion-App > ")
		reader.Scan()

		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]
		args := []string{}
		if len(words) > 1 {
			args = words[1:]
		}

		command, exists := getCommands()[commandName]
		if exists {
			err := command.callback(cfg, args...)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unkown command")
			continue
		}
	}
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays help message",
			callback:    commandHelp,
		},
		"books": {
			name:        "books",
			description: "Lists all LOTR books",
			callback:    commandGetBooks,
		},
		"characters": {
			name:        "characters",
			description: "Lists all characters",
			callback:    commandGetCharacters,
		},
		"movies": {
			name:        "movies",
			description: "List all LOTR movies",
			callback:    commandGetMovies,
		},
		"details": {
			name:        "details",
			description: "Return details about Movie or Book",
			callback:    commandGetDetails,
		},
		"quotes": {
			name:        "quotes",
			description: "List next page of quotes",
			callback:    commandQuotesf,
		},
		"quotesb": {
			name:        "quotesb",
			description: "Gets previous page of quotes",
			callback:    commandQuotesb,
		},
		"exit": {
			name:        "exit",
			description: "exit the program",
			callback:    commandExit,
		},
	}
}
