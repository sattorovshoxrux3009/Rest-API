package mysql

import (
	"context"
	"database/sql"
	"time"

	"example.com/m/storage/repo"
)

type postRepo struct {
	db *sql.DB
}

func NewPostStorage(db *sql.DB) repo.PostStorageI {
	return &postRepo{
		db: db,
	}
}

func (u *postRepo) Create(ctx context.Context, req *repo.Post) (*repo.Post, error) {
	query := `
		INSERT INTO posts (
			id, title,
			body, published,
			user_id
		) VALUES (?, ?, ?, ?, ?)
	`
	_, err := u.db.Exec(query, req.ID, req.Title, req.Body, true, req.UserId)
	if err != nil {
		return nil, err
	}
	// Agar req.CreateAt bo'sh bo'lsa, hozirgi vaqtni avtomatik ravishda olish
	if req.CreateAt.IsZero() {
		req.CreateAt = time.Now()
	}
	req.Published = true
	return req, nil
}

func (u *postRepo) Update(ctx context.Context, req *repo.UpdatePost) error {
	tsx, err := u.db.Begin()
	if err != nil {
		return err
	}
	query := `
		UPDATE posts SET 
			title=?,
			body=?,
			published=?
		WHERE id=?
	`
	res, err := tsx.Exec(query, req.Title, req.Body, req.Published, req.ID)
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

func (u *postRepo) Get(ctx context.Context, id string) (*repo.Post, error) {
	query := `
		SELECT 
			id, title,
			body, published, 
			user_id, created_at
		FROM posts WHERE id=?
	`
	var post repo.Post
	var createdAt []byte
	err := u.db.QueryRow(query, id).Scan(
		&post.ID,
		&post.Title,
		&post.Body,
		&post.Published,
		&post.UserId,
		&createdAt,
	)
	if err != nil {
		return nil, err
	}
	// Agar createdAt mavjud bo'lsa, uni time.Time formatiga o'zgartiramiz
	if len(createdAt) > 0 {
		parsedTime, err := time.Parse("2006-01-02 15:04:05", string(createdAt))
		if err != nil {
			return nil, err
		}
		post.CreateAt = parsedTime
	}
	return &post, nil
}

func (u *postRepo) Delete(ctx context.Context, id string) error {
	tsx, err := u.db.Begin()
	if err != nil {
		return err
	}
	res, err := tsx.Exec("DELETE FROM posts WHERE id=?", id)
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
