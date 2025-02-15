package theoneapi

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

func (c *Client) ListBooks() (booksResponse, error) {
	url := baseURL + "/book"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return booksResponse{}, errors.New("error")
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return booksResponse{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return booksResponse{}, err
	}

	booksResp := booksResponse{}
	err = json.Unmarshal(dat, &booksResp)
	if err != nil {
		return booksResponse{}, err
	}

	return booksResp, nil
}
