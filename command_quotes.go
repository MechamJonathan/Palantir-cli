package main

import (
	"errors"
	"fmt"
)

func commandQuotesf(cfg *config, args ...string) error {
	cfg.currentQuotePage += 1

	quotesResp, err := cfg.theoneapiClient.ListQuotes(cfg.currentQuotePage)
	if err != nil {
		return err
	}

	for _, quote := range quotesResp.Docs {
		fmt.Printf("=====================================\n")
		fmt.Printf("\"%s\"\n", quote.Dialog)
		fmt.Printf("- By %s\n", quote.CharacterName)
		fmt.Printf("=====================================\n\n")
	}

	return nil
}

func commandQuotesb(cfg *config, args ...string) error {

	if cfg.currentQuotePage <= 1 {
		return errors.New("you're on the first page")
	}

	cfg.currentQuotePage -= 1

	quotesResp, err := cfg.theoneapiClient.ListQuotes(cfg.currentQuotePage)
	if err != nil {
		return err
	}

	for _, quote := range quotesResp.Docs {
		fmt.Printf("=====================================\n")
		fmt.Printf("\"%s\"\n", quote.Dialog)
		fmt.Printf("- By %s\n", quote.CharacterName)
		fmt.Printf("=====================================\n\n")
	}

	return nil
}
