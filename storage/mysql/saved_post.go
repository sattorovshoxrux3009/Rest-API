package mysql

import (
	"context"
	"database/sql"

	"example.com/m/storage/repo"
)

type savedPostRepo struct {
	db *sql.DB
}

func NewSavedPostStorage(db *sql.DB) repo.SavedPostStorageI {
	return &savedPostRepo{
		db: db,
	}
}

func (u *savedPostRepo) Create(ctx context.Context, req *repo.SavedPost) (*repo.SavedPost, error) {
	query := `
		INSERT INTO saved_posts (
			id, post_id,
			user_id
		) VALUES (?, ?, ?)
	`
	_, err := u.db.Exec(query, req.ID, req.PostID, req.UserID)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func (u *savedPostRepo) Get(ctx context.Context, id string) (*repo.SavedPost, error) {
	query := `
		SELECT 
			id, post_id,
			user_id
		FROM saved_posts WHERE id=?
	`
	var savedPost repo.SavedPost
	err := u.db.QueryRow(query, id).Scan(
		&savedPost.ID,
		&savedPost.PostID,
		&savedPost.UserID,
	)
	if err != nil {
		return nil, err
	}
	return &savedPost, nil
}

func (u *savedPostRepo) Delete(ctx context.Context, id string) error {
	tsx, err := u.db.Begin()
	if err != nil {
		return err
	}
	res, err := tsx.Exec("DELETE FROM saved_posts WHERE id=?", id)
	if err != nil {
		errRoll := tsx.Rollback()
		if errRoll != nil {
			err = errRoll
		}
		return err
	}
	data, err := res.RowsAffected()
	if err != nil {
		errRoll := tsx.Rollback()
		if errRoll != nil {
			err = errRoll
		}
		return err
	}
	if data == 0 {
		tsx.Commit()
		return sql.ErrNoRows
	}
	return tsx.Commit()
}
