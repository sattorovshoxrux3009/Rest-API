package models

type CreateSavedPost struct {
	PostID string `json:"post_id"`
	UserID string `json:"user_id"`
}
type SavedPost struct {
	ID     string `json:"id"`
	PostID string `json:"post_id"`
	UserID string `json:"user_id"`
}
