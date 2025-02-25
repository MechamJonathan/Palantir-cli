package theoneapi

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strings"
)

func (c *Client) ListMovies() (MovieResponse, error) {
	url := baseURL + "/movie"

	if val, ok := c.cache.Get(url); ok {
		movieResp := MovieResponse{}
		err := json.Unmarshal(val, &movieResp)
		if err != nil {
			return MovieResponse{}, err
		}

		return movieResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return MovieResponse{}, err
	}

	req.Header.Set("Authorization", "Bearer "+c.apiKey)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return MovieResponse{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return MovieResponse{}, err
	}

	if len(dat) == 0 {
		return MovieResponse{}, errors.New("received empty response from API")
	}

	movieResp := MovieResponse{}
	err = json.Unmarshal(dat, &movieResp)
	if err != nil {
		return MovieResponse{}, err
	}

	c.cache.Add(url, dat)

	for _, movie := range movieResp.Docs {
		c.cache.Add("movie:"+strings.ToLower(movie.Name), []byte(movie.ID))
	}

	return movieResp, nil
}
