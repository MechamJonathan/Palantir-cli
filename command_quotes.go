package main

import (
	"errors"
	"fmt"
	"strings"
)

func commandQuotesf(cfg *config, args ...string) error {
	inputName := strings.Join(args, " ")

	cfg.currentQuotePage += 1

	quotesResp, err := cfg.theoneapiClient.ListQuotes(inputName, cfg.currentQuotePage)
	if err != nil {
		return err
	}

	for _, quote := range quotesResp.Docs {
		fmt.Printf("------------------------------------------\n")
		fmt.Printf("\"%s\"\n", quote.Dialog)
		fmt.Printf("- By %s\n", quote.CharacterName)
		fmt.Printf("------------------------------------------\n\n")
	}

	return nil
}

func commandQuotesb(cfg *config, args ...string) error {
	inputName := strings.Join(args, " ")

	if cfg.currentQuotePage <= 1 {
		return errors.New("you're on the first page of quotes")
	}

	cfg.currentQuotePage -= 1

	quotesResp, err := cfg.theoneapiClient.ListQuotes(inputName, cfg.currentQuotePage)
	if err != nil {
		return err
	}

	for _, quote := range quotesResp.Docs {
		fmt.Printf("------------------------------------------\n")
		fmt.Printf("\"%s\"\n", quote.Dialog)
		fmt.Printf("- By %s\n", quote.CharacterName)
		fmt.Printf("------------------------------------------\n\n")
	}

	return nil
}
