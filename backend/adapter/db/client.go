package db

import (
	"time"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm/logger"

	_ "github.com/lib/pq"
	"gorm.io/gorm"
)

var dbClient DbClient

func init() {
	initClient()
}

func initClient() {
	connStr := "host=localhost port=5432 user=admin dbname=airlineDb password=dbpassword sslmode=disable"
	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{
		Logger: logger.Discard,
	})
	if err != nil {
		logrus.Fatal(err)
	}

	postgresDb, err := db.DB()
	if err != nil {
		logrus.Fatal(err)
	}
	postgresDb.SetMaxIdleConns(5)
	postgresDb.SetMaxOpenConns(20)
	postgresDb.SetConnMaxLifetime(10 * time.Minute)
	dbClient = &pgConnection{db: db}
}

func GetClient() DbClient {
	return dbClient
}

func (pgc *pgConnection) GetDb() *gorm.DB {
	return pgc.db
}

func (pgc *pgConnection) Lock(lockId int) bool {
	// Obtain exclusive session level advisory lock
	// https://www.postgresql.org/docs/current/functions-admin.html#FUNCTIONS-ADVISORY-LOCKS
	var status LockStatus
	err := pgc.db.Raw("SELECT pg_try_advisory_lock(?)", lockId).First(&status).Error
	if err != nil {
		logrus.Error(err)
		return false
	}
	return status.Status
}

func (pgc *pgConnection) Unlock(lockId int) error {
	// Release the advisory lock
	err := pgc.db.Raw("SELECT pg_advisory_unlock(?)", lockId).Error
	if err != nil {
		logrus.Error(err)
	}
	return err
}
