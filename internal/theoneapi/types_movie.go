package theoneapi

type Movie struct {
	ID                         string  `json:"_id"`
	Name                       string  `json:"name"`
	RuntimeInMinutes           int     `json:"runtimeInMinutes"`
	BudgetInMillions           float64 `json:"budgetInMillions"`
	BoxOfficeRevenueInMillions float64 `json:"boxOfficeRevenueInMillions"`
	AcademyAwardNominations    int     `json:"academyAwardNominations"`
	AcademyAwardWins           int     `json:"academyAwardWins"`
	RottenTomatoesScore        float64 `json:"rottenTomatoesScore"`
}

type MovieResponse struct {
	Docs []Movie `json:"docs"`
}
