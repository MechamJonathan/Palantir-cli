package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/MechamJonathan/palantir-cli/internal/theoneapi"
	"github.com/MechamJonathan/palantir-cli/styles"
)

var startUpQuotes = []string{
	"“...Tʜᴇʏ ᴀʀᴇ ɴᴏᴛ ᴀʟʟ ᴀᴄᴄᴏᴜɴᴛᴇᴅ ғᴏʀ, ᴛʜᴇ ʟᴏsᴛ Sᴇᴇɪɴɢ Sᴛᴏɴᴇs.\n\n	Wᴇ ᴅᴏ ɴᴏᴛ ᴋɴᴏᴡ ᴡʜᴏ ᴇʟsᴇ ᴍᴀʏ ʙᴇ ᴡᴀᴛᴄʜɪɴɢ...”",
	"“A Pᴀʟᴀɴᴛɪ́ʀ ɪs ᴀ ᴅᴀɴɢᴇʀᴏᴜs ᴛᴏᴏʟ, Sᴀʀᴜᴍᴀɴ...\n\n	...Wʜʏ? Wʜʏ sʜᴏᴜʟᴅ ᴡᴇ ғᴇᴀʀ ᴛᴏ ᴜsᴇ ɪᴛ”",
	"“Dɪᴅ I ɴᴏᴛ ᴛᴇʟʟ ʏᴏᴜ, Pᴇʀᴇɢʀɪɴ Tᴏᴏᴋ, ɴᴇᴠᴇʀ ᴛᴏ ʜᴀɴᴅʟᴇ ɪᴛ?”"}

type config struct {
	theoneapiClient      theoneapi.Client
	currentQuotePage     int
	currentCharacterName string
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

func startRepl(cfg *config) {
	reader := bufio.NewScanner(os.Stdin)
	if err := ClearScreen(); err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	quote := getRandomQuote(startUpQuotes)
	fmt.Println(styles.StartUpQuote.Render(quote))
	MoveCursorToBottom()
	cfg.currentQuotePage = 0

	for {
		fmt.Print(styles.PalitirStyle.Render("Palantír"),
			styles.ArrowSymbol.Render(" > "))

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
			if err := ClearScreen(); err != nil {
				fmt.Printf("Error: %v\n", err)
			}
			err := command.callback(cfg, args...)
			if err != nil {
				fmt.Println(err)
			}
			MoveCursorToBottom()
			continue
		} else {
			if err := ClearScreen(); err != nil {
				fmt.Printf("Error: %v\n", err)
			}
			fmt.Println("Unkown command")
			MoveCursorToBottom()
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
			description: "Display help message and all available commands",
			callback:    commandHelp,
		},
		"books": {
			name:        "books",
			description: "Lists all books",
			callback:    commandGetBooks,
		},
		"characters": {
			name:        "characters",
			description: "Lists all characters or group of characters",
			callback:    commandGetCharacters,
		},
		"movies": {
			name:        "movies",
			description: "List all LOTR movies",
			callback:    commandGetMovies,
		},
		"details": {
			name:        "details",
			description: "Return details about specific character, movie, or book",
			callback:    commandGetDetails,
		},
		"quotes": {
			name:        "quotes",
			description: "View next page of a character's quotes",
			callback:    commandQuotesf,
		},
		"quotesb": {
			name:        "quotesb",
			description: "View previous page of a character's quotes  ",
			callback:    commandQuotesb,
		},
		"exit": {
			name:        "exit",
			description: "Exit the program",
			callback:    commandExit,
		},
	}
}
