package storage

import (
	"database/sql"

	"example.com/m/storage/mysql"
	"example.com/m/storage/repo"
)

type StorageI interface {
	User() repo.UserStorageI
	Post() repo.PostStorageI
}

type storagePg struct {
	userRepo repo.UserStorageI
	postRepo repo.PostStorageI
}

func NewStorage(mysqlConn *sql.DB) StorageI {
	return &storagePg{
		userRepo: mysql.NewUserStorage(mysqlConn),
		postRepo: mysql.NewPostStorage(mysqlConn),
	}
}

func (s *storagePg) User() repo.UserStorageI {
	return s.userRepo
}
func (s *storagePg) Post() repo.PostStorageI {
	return s.postRepo
}
