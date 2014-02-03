package ushahidi

// MessagesService provides access to the Messages Resource
//
// Ushahidi v3 API docs: https://wiki.ushahidi.com/display/WIKI/Ushahidi+3.x+API+Messages+Resource
type MessagesService struct {
	client *Client
}

// Message represents Ushahidi message.
type Message struct {
	Id int
}
