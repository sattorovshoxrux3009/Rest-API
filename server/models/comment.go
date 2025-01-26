package models

type CreateComment struct {
	Body   string `json:"body"`
	PostId string `json:"post_id"`
	UserId string `json:"user_id"`
}
type Comment struct {
	ID       string `json:"id"`
	Body     string `json:"body"`
	PostId   string `json:"post_id"`
	UserId   string `json:"user_id"`
	CreateAt string `json:"created_at"`
}
type UpdateComment struct {
	Body string `json:"body"`
}
