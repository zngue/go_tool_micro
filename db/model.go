package db

import "time"
type TimeStampModel struct {
	CreatedAt int64 `gorm:"column:created_at" json:"created_at" `
	UpdatedAt int64 `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt *time.Time `gorm:"column:deleted_at" json:"deleted_at" sql:"index"`
}
type GormModel struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

