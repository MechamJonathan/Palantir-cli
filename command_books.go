package main

import "fmt"

func commandGetBooks(cfg *config, args ...string) error {
	booksResp, err := cfg.theoneapiClient.ListBooks()
	if err != nil {
		return err
	}

	//fmt.Println("\n--- List of Books ---")
	fmt.Printf("\n%-20s\n", "Books")
	fmt.Println("--------------------")
	for _, book := range booksResp.Docs {
		fmt.Println(" -", book.Name)
	}
	fmt.Println("")
	return nil
}
