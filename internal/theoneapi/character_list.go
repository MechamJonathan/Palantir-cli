package theoneapi

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strings"
)

func (c *Client) ListCharacters() (CharacterResponse, error) {
	url := baseURL + "/character?sort=name:asc"

	if val, ok := c.cache.Get(url); ok {
		charResp := CharacterResponse{}
		err := json.Unmarshal(val, &charResp)
		if err != nil {
			return CharacterResponse{}, err
		}

		return charResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return CharacterResponse{}, err
	}

	req.Header.Set("Authorization", "Bearer "+c.apiKey)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return CharacterResponse{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return CharacterResponse{}, err
	}

	if len(dat) == 0 {
		return CharacterResponse{}, errors.New("received empty response from API")
	}

	charResp := CharacterResponse{}
	err = json.Unmarshal(dat, &charResp)
	if err != nil {
		return CharacterResponse{}, err
	}

	c.cache.Add(url, dat)

	for _, char := range charResp.Docs {
		c.cache.Add("character:"+strings.ToLower(char.Name), []byte(char.ID))
	}

	return charResp, nil
}
