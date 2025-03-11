package theoneapi

type Quote struct {
	ID            string `json:"_id"`
	Dialog        string `json:"dialog"`
	Movie         string `json:"movie"`
	Character     string `json:"character"`
	CharacterName string `json:"character_name,omitempty"` // New field for name
}

type QuoteResponse struct {
	Docs []Quote `json:"docs"`
}
