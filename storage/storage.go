package storage

import (
	"database/sql"

	"example.com/m/storage/mysql"
	"example.com/m/storage/repo"
)

type StorageI interface {
	User() repo.UserStorageI
	Post() repo.PostStorageI
	Comment() repo.CommentStorageI
	Saved_post() repo.SavedPostStorageI
}

type storagePg struct {
	userRepo      repo.UserStorageI
	postRepo      repo.PostStorageI
	commentRepo   repo.CommentStorageI
	savedPostRepo repo.SavedPostStorageI
}

func NewStorage(mysqlConn *sql.DB) StorageI {
	return &storagePg{
		userRepo:      mysql.NewUserStorage(mysqlConn),
		postRepo:      mysql.NewPostStorage(mysqlConn),
		commentRepo:   mysql.NewCommentStorage(mysqlConn),
		savedPostRepo: mysql.NewSavedPostStorage(mysqlConn),
	}
}

func (s *storagePg) User() repo.UserStorageI {
	return s.userRepo
}
func (s *storagePg) Post() repo.PostStorageI {
	return s.postRepo
}
func (s *storagePg) Comment() repo.CommentStorageI {
	return s.commentRepo
}
func (s *storagePg) Saved_post() repo.SavedPostStorageI {
	return s.savedPostRepo
}
