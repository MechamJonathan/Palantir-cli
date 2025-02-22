package theoneapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListQuotes(nextPage int) (QuoteResponse, error) {
	url := fmt.Sprintf("%s/quote?limit=20&page=%d", baseURL, nextPage)

	if val, ok := c.cache.Get(url); ok {
		quoteResp := QuoteResponse{}
		err := json.Unmarshal(val, &quoteResp)
		if err != nil {
			return QuoteResponse{}, err
		}

		return quoteResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return QuoteResponse{}, err
	}

	req.Header.Set("Authorization", "Bearer "+c.apiKey)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return QuoteResponse{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return QuoteResponse{}, err
	}

	quoteResp := QuoteResponse{}
	err = json.Unmarshal(dat, &quoteResp)
	if err != nil {
		return QuoteResponse{}, err
	}

	for i, quote := range quoteResp.Docs {
		if name, ok := c.cache.Get("character:" + quote.Character); ok {
			quoteResp.Docs[i].CharacterName = string(name)
		} else {
			character, err := c.FetchCharacterByID(quote.Character)
			if err == nil {
				quoteResp.Docs[i].CharacterName = character.Name
				c.cache.Add("character:"+quote.Character, []byte(character.Name))
			} else {
				quoteResp.Docs[i].CharacterName = "Unknown Character"
			}
		}
	}

	enhancedData, _ := json.Marshal(quoteResp)
	c.cache.Add(url, enhancedData)

	return quoteResp, nil
}
