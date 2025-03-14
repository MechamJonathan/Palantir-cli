package theoneapi

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strings"
)

func (c *Client) GetBookByName(bookName string) (Book, error) {
	if val, ok := c.cache.Get("book:" + strings.ToLower(bookName)); ok {
		bookID := string(val)
		return c.fetchBookByID(bookID)
	}

	bookResp, err := c.ListBooks()
	if err != nil {
		return Book{}, err
	}

	for _, book := range bookResp.Docs {
		if strings.ToLower(book.Name) == bookName {
			c.cache.Add("book:"+strings.ToLower(book.Name), []byte(book.ID))
			return c.fetchBookByID(book.ID)
		}
	}

	return Book{}, errors.New("Book not found")
}

func (c *Client) fetchBookByID(bookID string) (Book, error) {
	url := baseURL + "/book/" + bookID

	if val, ok := c.cache.Get(url); ok {
		bookResp := BooksResponse{}
		err := json.Unmarshal(val, &bookResp)
		if err != nil {
			return Book{}, err
		}
		if len(bookResp.Docs) > 0 {
			return bookResp.Docs[0], nil
		}

		return Book{}, errors.New("Book not found in cached data")
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Book{}, err
	}
	req.Header.Set("Authorization", "Bearer "+c.apiKey)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Book{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return Book{}, err
	}

	if len(dat) == 0 {
		return Book{}, errors.New("received empty response from API")
	}

	bookResp := BooksResponse{}
	err = json.Unmarshal(dat, &bookResp)
	if err != nil {
		return Book{}, err
	}

	if len(bookResp.Docs) > 0 {
		return bookResp.Docs[0], nil
	}

	return Book{}, errors.New("Book not found in API response")
}
