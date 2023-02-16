package models

import "time"

type Category struct {
	ID          uint      `gorm:"primaryKey;autoIncrement;column:id_category" json:"id_category"`
	Name        string    `gorm:"type:varchar(150);column:name" json:"name"`
	Description string    `gorm:"column:description" json:"description"`
	Deleted     string    `gorm:"type:varchar(10);column:deleted;default:'N'" json:"deleted"`
	IDCreated   uint      `gorm:"column:id_created" json:"id_created"`
	IDUpdated   uint      `gorm:"column:id_updated" json:"id_updated"`
	CreatedAt   time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (Category) TableName() string {
	return "master_categories"
}
