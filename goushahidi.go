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
	"net/http"
	"net/url"
)

const (
	libraryVersion = "0.1"
	defaultBaseURL = "http://demo.ushahidi.com" // Demo Ushahidi Instance
)

type UshahidiClient struct {
	client  *http.Client
	BaseURL *url.URL
}

func NewClient() *UshahidiClient {
	baseURL, _ := url.Parse(defaultBaseURL)
	return &UshahidiClient{client: http.DefaultClient, BaseURL: baseURL}

}

func (u *UshahidiClient) NewRequest(method, urlStr string, body interface{}) {

}
