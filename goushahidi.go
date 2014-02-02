/*
Package goushahidi provides a client for accessing data from Ushahidi API.

Access different data using the API:
         client := goushahidi.NewClient(nil)

         // list all incidences
         incidences, err := client.GetIncidences()


The full Ushahidi API is documented at https://wiki.ushahidi.com/display/WIKI/Ushahidi+3.x+REST+API.
*/
package goushahidi

import (
	"bytes"
	"encoding/json"
	"flag"
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
}

func NewClient(DefaultBaseURL string) *Client {
	flag.Parse()
	baseURL, _ := url.Parse(DefaultBaseURL)
	return &Client{client: http.DefaultClient, BaseURL: baseURL}

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
