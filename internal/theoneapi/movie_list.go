package theoneapi

import (
	"encoding/json"
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

	movieResp := MovieResponse{}
	err = json.Unmarshal(dat, &movieResp)
	if err != nil {
		return MovieResponse{}, err
	}

	// Cache full response
	c.cache.Add(url, dat)

	// Cache individual movie names -> IDs
	for _, movie := range movieResp.Docs {
		c.cache.Add("movie:"+strings.ToLower(movie.Name), []byte(movie.ID))
	}

	return movieResp, nil
}
