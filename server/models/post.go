package models

type CreatePost struct {
	Title     string `json:"title"`
	Body      string `json:"body"`
	UserId    string `json:"user_id"`
}
type Post struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	Body      string `json:"body"`
	Published bool   `json:"published"`
	UserId    string `json:"user_id"`
	CreateAt  string `json:"created_at"`
}
type UpdatePost struct {
	Title     string `json:"title"`
	Body      string `json:"body"`
	Published bool   `json:"published"`
}
