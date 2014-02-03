package ushahidi

import (
	"fmt"
)

// UsersService provides access to the Users Resource
//
// Ushahidi v3 API docs: https://wiki.ushahidi.com/display/WIKI/Ushahidi+3.x+API+Users+Resource
type UsersService struct {
	client *Client
}

// User represents Ushahidi instance user.

type User struct {
	Id             int    `json:"id,omitempty"`
	Url            string `json:"url,omitempty"`
	Email          string `json:"email,omitempty"`
	FirstName      string `json:"first_name,omitempty"`
	LastName       string `json:"last_name,omitempty"`
	Username       string `json:"username,omitempty"`
	Logins         int    `json:"logins,omitempty"`
	LastLogin      string `json:"last_login,omitempty"`
	FailedAttempts int    `json:"failed_attempts,omitempty"`
	LastAttempt    string `json:"last_attempt,omitempty"`
	Created        string `json:"created,omitempty"`
	Updated        string `json:"updated,omitempty"`
}

// Get a single user.
//
// Ushahidi API docs: https://wiki.ushahidi.com/display/WIKI/Ushahidi+3.x+API+Users+Resource#Ushahidi3.xAPIUsersResource-GETusers/:id
func (u *UsersService) Get(id int) (*User, error) {
	url := fmt.Sprintf("users/%v", id)
	req, err := u.client.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	uResp := new(User)

}

// Create a new user for the specified instance.
//
// Ushahidi API docs: https://wiki.ushahidi.com/display/WIKI/Ushahidi+3.x+API+Users+Resource#Ushahidi3.xAPIUsersResource-POSTusers
func (u *UsersService) Create(email string, firstName string,
	lastName string, userName string, password string) (*User, error) {

}
