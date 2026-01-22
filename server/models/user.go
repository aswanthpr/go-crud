package models

import "time"

type User struct {
	ID        uint       `gorm:"primaryKey;autoIncrement;column:id"`
	Name      string     `gorm:"type:varchar(255);not null;column:name"`
	Email     string     `gorm:"type:varchar(255);not null;uniqueIndex;column:email"`
	Password   string     `gorm:"type:varchar(255);not null;column:password"`
	CreatedAt time.Time  `gorm:"autoCreateTime;column:created_at"`
	UpdatedAt *time.Time `gorm:"autoUpdateTime;column:updated_at"`
}

func (User) TableName() string {
	return "users"
}
