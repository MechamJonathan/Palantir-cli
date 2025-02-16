package theoneapi

type Movie struct {
	ID   string `json:"_id"`
	Name string `json:"name"`
}

type MovieResponse struct {
	Docs []Book `json:"docs"`
}
