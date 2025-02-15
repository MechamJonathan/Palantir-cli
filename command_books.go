package main

import "fmt"

func getBooks(cfg *config) error {
	booksResp, err := cfg.theoneapiClient.ListBooks()
	if err != nil {
		return err
	}

	fmt.Println("\n--List of Books--")
	for _, book := range booksResp.Docs {
		fmt.Println(" -", book.Name)
	}
	fmt.Println("")
	return nil
}
