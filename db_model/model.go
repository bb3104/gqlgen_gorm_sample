package db_model

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

type User struct {
	ID        uint32    `gorm:"primary_key;AUTO_INCREMENT;"`
	Name      string    `gorm:"size:255"`
	Password  string    `gorm:"size:255"`
	Email     string    `gorm:"size:255"`
	CreatedAt time.Time `gorm:"column:created_at;type:datetime;"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:datetime;"`
	Items     []Item    `gorm:"foreignkey:ID"`
}

type Item struct {
	ID        uint32 `gorm:"primary_key;AUTO_INCREMENT;"`
	Name      string
	UserID    uint32
	CreatedAt time.Time `gorm:"column:created_at;type:datetime;"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:datetime;"`
}
