package storage

import "database/sql"

type StorageI interface {
}

type StoragePg struct {
}

func NewStorage(mysqlConn *sql.DB) StorageI {

	return &StoragePg{
		
	}
}
