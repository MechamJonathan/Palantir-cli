package theoneapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListMovies() (MovieResponse, error) {
	url := baseURL + "/movie"

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

	return movieResp, nil
}
