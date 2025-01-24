package mysql

import (
	"context"
	"database/sql"
	"time"

	"example.com/m/storage/repo"
)

type userRepo struct {
	db *sql.DB
}

func NewUserStorage(db *sql.DB) repo.UserStorageI {
	return &userRepo{
		db: db,
	}
}

func (u *userRepo) Create(ctx context.Context, req *repo.User) (*repo.User, error) {
	query := `
		INSERT INTO users (
			id, first_name,
			last_name, email,
			password
		) VALUES (?, ?, ?, ?, ?)
	`
	_, err := u.db.Exec(query, req.ID, req.FirstName, req.LastName, req.Email, req.Password)
	if err != nil {
		return nil, err
	}
	// Agar req.CreateAt bo'sh bo'lsa, hozirgi vaqtni avtomatik ravishda olish
	if req.CreateAt.IsZero() {
		req.CreateAt = time.Now()
	}
	return req, nil
}

func (u *userRepo) Update(ctx context.Context, req *repo.UpdateUser) error {
	tsx, err := u.db.Begin()
	if err != nil {
		return err
	}
	query := `
		UPDATE users SET 
			first_name=?,
			last_name=?
		WHERE id=?
	`
	res, err := tsx.Exec(query, req.FirstName, req.LastName, req.ID)
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

func (u *userRepo) Get(ctx context.Context, id string) (*repo.User, error) {
	query := `
		SELECT 
			id, first_name,
			last_name, email, 
			password, created_at
		FROM users WHERE id=?
	`
	var user repo.User
	var createdAt []byte
	err := u.db.QueryRow(query, id).Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Password,
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
		user.CreateAt = parsedTime
	}
	return &user, nil
}

func (u *userRepo) Delete(ctx context.Context, id string) error {
	tsx, err := u.db.Begin()
	if err != nil {
		return err
	}
	res, err := tsx.Exec("DELETE FROM users WHERE id=?", id)
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
