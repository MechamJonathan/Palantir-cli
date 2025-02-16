package theoneapi

type Book struct {
	ID   string `json:"_id"`
	Name string `json:"name"`
}

type BooksResponse struct {
	Docs []Book `json:"docs"`
}
