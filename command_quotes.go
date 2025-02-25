package main

import (
	"errors"
	"fmt"
	"strings"

	"github.com/MechamJonathan/lotr-companion-app/internal/theoneapi"
)

func commandQuotesf(cfg *config, args ...string) error {
	inputName := strings.Join(args, " ")

	if len(args) < 1 && cfg.currentCharacterName == "" {
		return errors.New("usage: quotes <character name>")
	} else if cfg.currentCharacterName == "" {
		cfg.currentCharacterName = inputName
	} else if cfg.currentCharacterName != "" && len(args) >= 1 {
		cfg.currentCharacterName = inputName
		cfg.currentQuotePage = 1
	}

	cfg.currentQuotePage += 1

	quotesResp, err := cfg.theoneapiClient.ListQuotes(cfg.currentCharacterName, cfg.currentQuotePage)
	if err != nil {
		return err
	}

	printQuotes(quotesResp.Docs)

	return nil
}

func commandQuotesb(cfg *config, args ...string) error {
	if cfg.currentQuotePage <= 1 {
		return errors.New("you're on the first page of quotes")
	}

	cfg.currentQuotePage -= 1

	quotesResp, err := cfg.theoneapiClient.ListQuotes(cfg.currentCharacterName, cfg.currentQuotePage)
	if err != nil {
		return err
	}

	printQuotes(quotesResp.Docs)

	return nil
}

func printQuotes(quotes []theoneapi.Quote) {
	for _, quote := range quotes {
		fmt.Printf("------------------------------------------\n")
		fmt.Printf("\"%s\"\n", quote.Dialog)
		fmt.Printf("- %s\n", quote.CharacterName)
		fmt.Printf("------------------------------------------\n\n")
	}
}
