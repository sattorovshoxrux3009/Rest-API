package storage

import (
	"database/sql"

	"example.com/m/storage/mysql"
	"example.com/m/storage/repo"
)

type StorageI interface {
	User() repo.UserStorageI
}

type storagePg struct {
	userRepo repo.UserStorageI
}

func NewStorage(mysqlConn *sql.DB) StorageI {
	return &storagePg{
		userRepo: mysql.NewUserStorage(mysqlConn),
	}
}
func (s *storagePg) User() repo.UserStorageI {
	return s.userRepo
}
