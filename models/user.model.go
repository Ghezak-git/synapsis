package models

import "time"

type User struct {
	ID        uint      `gorm:"primaryKey;autoIncrement;column:id_user" json:"id_user"`
	Email     string    `gorm:"type:varchar(150);uniqueIndex;not null;column:email" json:"email"`
	FullName  string    `gorm:"type:varchar(150);column:fullname" json:"fullname"`
	Password  string    `gorm:"type:varchar(255);column:password" json:"password"`
	Role      string    `gorm:"type:varchar(20);column:role;default:'customer'" json:"role"`
	Status    string    `gorm:"type:varchar(20);column:status;default:'Active'" json:"status"`
	Deleted   string    `gorm:"type:varchar(10);column:deleted;default:'N'" json:"deleted"`
	IDCreated uint      `gorm:"column:id_created" json:"id_created"`
	IDUpdated uint      `gorm:"column:id_updated" json:"id_updated"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (User) TableName() string {
	return "master_users"
}

type LogInInput struct {
	Email    string `json:"email"  binding:"required"`
	Password string `json:"password"  binding:"required"`
}

type SignUpInput struct {
	Email           string `json:"email"  binding:"required"`
	FullName        string `json:"fullname"  binding:"required"`
	Password        string `json:"password" binding:"required,min=8"`
	PasswordConfirm string `json:"passwordConfirm" binding:"required"`
}
