package theoneapi

type Character struct {
	ID      string `json:"_id"`
	Name    string `json:"name"`
	WikiURL string `json:"wikiUrl"`
	Race    string `json:"race"`
	Birth   string `json:"birth"`
	Gender  string `json:"gender"`
	Death   string `json:"death"`
	Hair    string `json:"hair"`
	Height  string `json:"height"`
	Realm   string `json:"realm"`
	Spouse  string `json:"spouse"`
}

type CharacterResponse struct {
	Docs []Character `json:"docs"`
}
