package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"

	"github.com/MechamJonathan/lotr-companion-app/internal/theoneapi"
	"github.com/MechamJonathan/lotr-companion-app/styles"
)

var startUpQuotes = []string{"“...Tʜᴇʏ ᴀʀᴇ ɴᴏᴛ ᴀʟʟ ᴀᴄᴄᴏᴜɴᴛᴇᴅ ғᴏʀ, ᴛʜᴇ ʟᴏsᴛ Sᴇᴇɪɴɢ Sᴛᴏɴᴇs. Wᴇ ᴅᴏ ɴᴏᴛ ᴋɴᴏᴡ ᴡʜᴏ ᴇʟsᴇ ᴍᴀʏ ʙᴇ ᴡᴀᴛᴄʜɪɴɢ”",
	"“𝘈 𝘗𝘢𝘭𝘢𝘯𝘵𝘪́𝘳 𝘪𝘴 𝘢 𝘥𝘢𝘯𝘨𝘦𝘳𝘰𝘶𝘴 𝘵𝘰𝘰𝘭, 𝘚𝘢𝘳𝘶𝘮𝘢𝘯...\n\n   ...𝘞𝘩𝘺? 𝘞𝘩𝘺 𝘴𝘩𝘰𝘶𝘭𝘥 𝘸𝘦 𝘧𝘦𝘢𝘳 𝘵𝘰 𝘶𝘴𝘦 𝘪𝘵?”",
	"“𝐴𝑙𝑎𝑠! 𝑇ℎ𝑎𝑡 𝑡ℎ𝑖𝑛𝑔 𝑖𝑠 𝑏𝑒𝑦𝑜𝑛𝑑 𝑎𝑙𝑙 𝑜𝑓 𝑢𝑠 𝑒𝑥𝑐𝑒𝑝𝑡 𝑝𝑒𝑟ℎ𝑎𝑝𝑠 𝐴𝑟𝑎𝑔𝑜𝑟𝑛. 𝐷𝑖𝑑 𝐼 𝑛𝑜𝑡 𝑡𝑒𝑙𝑙 𝑦𝑜𝑢, 𝑃𝑒𝑟𝑒𝑔𝑟𝑖𝑛 𝑇𝑜𝑜𝑘, 𝑛𝑒𝑣𝑒𝑟 𝑡𝑜 ℎ𝑎𝑛𝑑𝑙𝑒 𝑖𝑡?”"}

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

func randomStartupQuote(quotes []string) (string, error) {
	randomIndex := rand.Intn(len(quotes))
	return quotes[randomIndex], nil
}

func startRepl(cfg *config) {
	reader := bufio.NewScanner(os.Stdin)
	if err := clearScreen(); err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	quote, _ := randomStartupQuote(startUpQuotes)
	fmt.Println(styles.SubHeader.Render(quote))
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
			err := command.callback(cfg, args...)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			if err := clearScreen(); err != nil {
				fmt.Printf("Error: %v\n", err)
			}
			fmt.Println(styles.ErrorMessage.Render("Unkown command"))
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
