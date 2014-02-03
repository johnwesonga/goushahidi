package ushahidi

// UsersService provides access to the Posts Resource
//
// Ushahidi v3 API docs: https://wiki.ushahidi.com/display/WIKI/Ushahidi+3.x+API+Posts+Resource
type PostsService struct {
	client *Client
}

// Post represents Ushahidi instance post.
type Post struct {
	Id   int    `json:"id,omitempty"`
	Url  string `json:"url,omitempty"`
	Form struct {
		Id  int    `json:"id,omitempty"`
		Url string `json:"url,omitempty"`
	}
	Locale  string `json:"locale,omitempty"`
	Type    string `json:"type,omitempty"`
	Title   string `json:"title,omitempty"`
	Status  string `json:"status,omitempty"`
	Content string `json:"content,omitempty"`
	Values  struct {
		FullName     string `json:"full_name,omitempty"`
		Description  string `json:"description,omitempty"`
		Dob          string `json:"dob,omitempty"`
		MissingDate  string `json:"missing_date,omitempty"`
		LastLocation string `json:"last_location,omitempty"`
		Status       string `json:"status,omitempty"`
	}
	Tags []string `json:"tags,omitempty"`
}

// Get a single Post.
//
// Ushahidi API docs: https://wiki.ushahidi.com/display/WIKI/Ushahidi+3.x+API+Posts+Resource#Ushahidi3.xAPIPostsResource-GETposts/:id
func (p *PostsService) Get(id int) (*Post, error) {

}

// Create a new post for the specified instance.
//
// Ushahidi API docs: https://wiki.ushahidi.com/display/WIKI/Ushahidi+3.x+API+Posts+Resource#Ushahidi3.xAPIPostsResource-POSTposts
func (p *PostsService) Create() (*Post, error) {

}
