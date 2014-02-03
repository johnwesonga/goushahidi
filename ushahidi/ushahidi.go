/*
Package ushahidi provides a client for accessing data from Ushahidi API.

Access different data using the API:
         client := ushahidi.NewClient(nil)

         // list all incidences
         incidences, err := client.GetIncidences()


The full Ushahidi API is documented at https://wiki.ushahidi.com/display/WIKI/Ushahidi+3.x+REST+API.
*/
package ushahidi

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/url"
)

const (
	libraryVersion = "0.1"
)

type Client struct {
	client         *http.Client
	BaseURL        *url.URL
	DefaultBaseURL string
	Users          *UsersService
	Posts          *PostsService
	Tags           *TagsService
	Messages       *MessagesService
	Media          *MediaService
}

func NewClient(DefaultBaseURL string, httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	baseURL, _ := url.Parse(DefaultBaseURL)
	c := &Client{client: httpClient, BaseURL: baseURL}
	c.Users = &UsersService{client: c}
	c.Posts = &PostsService{client: c}
	c.Tags = &TagsService{client: c}
	c.Media = &MediaService{client: c}
	c.Messages = &MessagesService{client: c}

	return c

}

func (c *Client) NewRequest(method, urlStr string, body interface{}) (*http.Request, error) {
	rel, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	u := c.BaseURL.ResolveReference(rel)
	buf := new(bytes.Buffer)
	if body != nil {
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}

	}

	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	return req, nil

}

// Do sends an API request and returns the API response.  The API response is
// decoded and stored in the value pointed to by v, or returned as an error if
// an API error has occurred.
func (c *Client) Do(req *http.Request, v *interface{}) (*http.Response, error) {
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	return resp, error

}

type Response struct {
	Limit  int
	Offset int
	Order  string
}
