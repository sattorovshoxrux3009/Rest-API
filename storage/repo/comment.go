package repo

import (
	"context"
	"time"
)

type CommentStorageI interface {
	Create(ctx context.Context, req *Comment) (*Comment, error)
	Get(ctx context.Context, id string) (*Comment, error)
	Update(ctx context.Context, req *UpdateComment) error
	Delete(ctx context.Context, id string) error
}

type Comment struct {
	ID       string
	Body     string
	PostId   string
	UserId   string
	CreateAt time.Time
}
type UpdateComment struct {
	ID        string
	Body      string
}
