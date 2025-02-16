package theoneapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListBooks() (BooksResponse, error) {
	url := baseURL + "/book"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return BooksResponse{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return BooksResponse{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return BooksResponse{}, err
	}

	booksResp := BooksResponse{}
	err = json.Unmarshal(dat, &booksResp)
	if err != nil {
		return BooksResponse{}, err
	}

	return booksResp, nil
}
