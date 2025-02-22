package theoneapi

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strings"
	"unicode"

	"golang.org/x/text/unicode/norm"
)

func (c *Client) GetCharacterByName(characterName string) (Character, error) {
	characterNameLower := strings.ToLower(characterName)

	if val, ok := c.cache.Get("character:" + characterNameLower); ok {
		characterID := string(val)
		return c.FetchCharacterByID(characterID)
	}

	charResp, err := c.ListCharacters()
	if err != nil {
		return Character{}, err
	}

	for _, character := range charResp.Docs {
		if strings.ToLower(removeDiacritics(character.Name)) == characterNameLower {
			c.cache.Add("character:"+strings.ToLower(removeDiacritics(character.Name)), []byte(character.ID))
			return c.FetchCharacterByID(character.ID)
		}
	}

	return Character{}, errors.New("Character not found")
}

func removeDiacritics(s string) string {
	s = norm.NFD.String(s)
	result := []rune{}
	for _, r := range s {
		if unicode.Is(unicode.Mn, r) {
			continue
		}
		result = append(result, r)
	}
	return string(result)
}

func (c *Client) FetchCharacterByID(characterID string) (Character, error) {
	url := baseURL + "/character/" + characterID

	if val, ok := c.cache.Get(url); ok {
		charResp := CharacterResponse{}
		err := json.Unmarshal(val, &charResp)
		if err != nil {
			return Character{}, err
		}
		if len(charResp.Docs) > 0 {
			return charResp.Docs[0], nil
		}
		return Character{}, errors.New("Character not found in cached data")
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Character{}, err
	}
	req.Header.Set("Authorization", "Bearer "+c.apiKey)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Character{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return Character{}, err
	}

	if len(dat) == 0 {
		return Character{}, errors.New("received empty response from API")
	}

	charResp := CharacterResponse{}
	err = json.Unmarshal(dat, &charResp)
	if err != nil {
		return Character{}, err
	}

	if len(charResp.Docs) > 0 {
		return charResp.Docs[0], nil
	}

	return Character{}, errors.New("Character not found in API response")
}
