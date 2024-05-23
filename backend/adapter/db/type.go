package db

import "gorm.io/gorm"

type pgConnection struct {
	db *gorm.DB
}

type DbClient interface {
	GetDb() *gorm.DB
	Lock(lockId int) bool
	Unlock(lockId int) error
}

type LockStatus struct {
	Status bool `gorm:"column:pg_try_advisory_lock"`
}
