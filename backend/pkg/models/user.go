package models

type User struct {
	ID   uint   `gorm:"not null; column:id" json:"id"`
	Name string `gorm:"not null; column:name" json:"name"`
}
