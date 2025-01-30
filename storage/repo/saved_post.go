package repo

import (
	"context"
)

type SavedPostStorageI interface {
	Create(ctx context.Context, req *SavedPost) (*SavedPost, error)
	Get(ctx context.Context, id string) (*SavedPost, error)
	Delete(ctx context.Context, id string) error
}

type SavedPost struct {
	ID     string
	PostID string
	UserID string
}
