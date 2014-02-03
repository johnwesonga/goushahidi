package ushahidi

// MediaService provides access to the Media Resource.
//
// Ushahidi v3 API docs: https://wiki.ushahidi.com/display/WIKI/Ushahidi+3.x+API+Media+Resource
type MediaService struct {
	client *Client
}

// Message represents Ushahidi message.
type Media struct {
	Id int
}
