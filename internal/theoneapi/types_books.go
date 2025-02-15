package theoneapi

type RespShallowBooks struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

type Book struct {
	ID   string `json:"_id"`
	Name string `json:"name"`
}

type booksResponse struct {
	Docs []Book `json:"docs"`
}
