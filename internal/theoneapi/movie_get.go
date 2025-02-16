package theoneapi

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strings"
)

func (c *Client) GetMovieByName(movieName string) (Movie, error) {
	if val, ok := c.cache.Get("movie:" + strings.ToLower(movieName)); ok {
		movieID := string(val)
		return c.fetchMovieByID(movieID)
	}

	movieResp, err := c.ListMovies()
	if err != nil {
		return Movie{}, err
	}

	for _, movie := range movieResp.Docs {
		if strings.ToLower(movie.Name) == movieName {
			c.cache.Add("movie:"+strings.ToLower(movie.Name), []byte(movie.ID))
			return c.fetchMovieByID(movie.ID)
		}
	}

	return Movie{}, errors.New("movie not found")
}

func (c *Client) fetchMovieByID(movieID string) (Movie, error) {
	url := baseURL + "/movie/" + movieID

	if val, ok := c.cache.Get(url); ok {
		movieResp := MovieResponse{}
		err := json.Unmarshal(val, &movieResp)
		if err != nil {
			return Movie{}, err
		}
		if len(movieResp.Docs) > 0 {
			return movieResp.Docs[0], nil
		}
		return Movie{}, errors.New("movie not found in cached data")
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Movie{}, err
	}
	req.Header.Set("Authorization", "Bearer "+c.apiKey)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Movie{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return Movie{}, err
	}

	if len(dat) == 0 {
		return Movie{}, errors.New("received empty response from API")
	}

	movieResp := MovieResponse{}
	err = json.Unmarshal(dat, &movieResp)
	if err != nil {
		return Movie{}, err
	}

	if len(movieResp.Docs) > 0 {
		return movieResp.Docs[0], nil
	}

	return Movie{}, errors.New("movie not found in API response")
}
