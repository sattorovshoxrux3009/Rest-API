package repo

import (
	"context"
	"time"
)

type PostStorageI interface {
	Create(ctx context.Context, req *Post) (*Post, error)
	Get(ctx context.Context, id string) (*Post, error)
	Update(ctx context.Context, req *UpdatePost) error
	Delete(ctx context.Context, id string) error
}

type Post struct {
	ID        string
	Title     string
	Body      string
	Published bool
	UserId    string
	CreateAt  time.Time
}
type UpdatePost struct {
	ID        string
	Title     string
	Body      string
	Published bool
}
