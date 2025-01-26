package mysql

import (
	"context"
	"database/sql"
	"time"

	"example.com/m/storage/repo"
)

type commentRepo struct {
	db *sql.DB
}

func NewCommentStorage(db *sql.DB) repo.CommentStorageI {
	return &commentRepo{
		db: db,
	}
}

func (u *commentRepo) Create(ctx context.Context, req *repo.Comment) (*repo.Comment, error) {
	query := `
		INSERT INTO comments (
			id, body,
			post_id, user_id
		) VALUES (?, ?, ?, ?)
	`
	_, err := u.db.Exec(query, req.ID, req.Body, req.PostId, req.UserId)
	if err != nil {
		return nil, err
	}
	// Agar req.CreateAt bo'sh bo'lsa, hozirgi vaqtni avtomatik ravishda olish
	if req.CreateAt.IsZero() {
		req.CreateAt = time.Now()
	}
	return req, nil
}

func (u *commentRepo) Update(ctx context.Context, req *repo.UpdateComment) error {
	tsx, err := u.db.Begin()
	if err != nil {
		return err
	}
	query := `
		UPDATE comments SET 
			body=?
		WHERE id=?
	`
	res, err := tsx.Exec(query, req.Body, req.ID)
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

func (u *commentRepo) Get(ctx context.Context, id string) (*repo.Comment, error) {
	query := `
		SELECT 
			id, body,
			post_id, user_id,
			created_at
		FROM comments WHERE id=?
	`
	var comment repo.Comment
	var createdAt []byte
	err := u.db.QueryRow(query, id).Scan(
		&comment.ID,
		&comment.Body,
		&comment.PostId,
		&comment.UserId,
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
		comment.CreateAt = parsedTime
	}
	return &comment, nil
}

func (u *commentRepo) Delete(ctx context.Context, id string) error {
	tsx, err := u.db.Begin()
	if err != nil {
		return err
	}
	res, err := tsx.Exec("DELETE FROM comments WHERE id=?", id)
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
