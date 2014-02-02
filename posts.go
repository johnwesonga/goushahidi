package goushahidi

// UsersService provides access to the Posts Resource
//
// Ushahidi v3 API docs: https://wiki.ushahidi.com/display/WIKI/Ushahidi+3.x+API+Posts+Resource
type PostsService struct {
	client *Client
}

type Post struct {
	Count   int
	Results *[]Results
}

type Results struct {
	Id   int
	Url  string
	Form struct {
		Id  int
		Url string
	}
	Locale  string
	Type    string
	Title   string
	Status  string
	Content string
	Values  struct {
		FullName     string
		Description  string
		Dob          string
		MissingDate  string
		LastLocation string
		Status       string
	}
	Tags []string
}
