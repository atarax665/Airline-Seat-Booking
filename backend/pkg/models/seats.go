package models

import "database/sql"

type Seat struct {
	ID     uint          `gorm:"not null; column:id" json:"id"`
	Name   string        `gorm:"not null; column:name" json:"name"`
	TripID uint          `gorm:"not null; column:trip_id" json:"trip_id"`
	UserID sql.NullInt64 `gorm:"column:user_id" json:"user_id"`
}
