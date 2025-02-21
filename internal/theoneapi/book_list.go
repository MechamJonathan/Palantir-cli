package theoneapi

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

func (c *Client) ListBooks() (BooksResponse, error) {
	url := baseURL + "/book"

	if val, ok := c.cache.Get(url); ok {
		movieResp := BooksResponse{}
		err := json.Unmarshal(val, &movieResp)
		if err != nil {
			return BooksResponse{}, err
		}

		return movieResp, nil
	}

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

	// Cache full response
	c.cache.Add(url, dat)

	// Cache individual movie names -> IDs
	for _, book := range booksResp.Docs {
		c.cache.Add("book:"+strings.ToLower(book.Name), []byte(book.ID))
	}

	return booksResp, nil
}
