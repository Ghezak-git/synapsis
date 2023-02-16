package models

import "time"

// Product
type Product struct {
	ID          uint      `gorm:"primaryKey;autoIncrement;column:id_product" json:"id_product"`
	IDCategory  uint      `gorm:"not null;column:id_category" json:"id_category"`
	Name        string    `gorm:"type:varchar(150);column:name" json:"name"`
	Description string    `gorm:"column:description" json:"description"`
	Price       int       `gorm:"column:price" json:"price"`
	Deleted     string    `gorm:"type:varchar(10);column:deleted;default:'N'" json:"deleted"`
	IDCreated   uint      `gorm:"column:id_created" json:"id_created"`
	IDUpdated   uint      `gorm:"column:id_updated" json:"id_updated"`
	CreatedAt   time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at" json:"updated_at"`
	Category    Category  `gorm:"foreignKey:id_category;references:id_category" json:"detail_category"`
}

func (Product) TableName() string {
	return "master_products"
}

type GetProductsRequest struct {
	Page       string `json:"page,omitempty" `
	PerPage    string `json:"perPage,omitempty"`
	Search     string `json:"search,omitempty"`
	Sort       string `json:"sort,omitempty"`
	SortColumn string `json:"sortColumn,omitempty"`
	Category   string `json:"category,omitempty"`
}

// Shop Cart
type ShopCart struct {
	IDProduct uint    `gorm:"primaryKey;column:id_product" json:"id_product"`
	IDUser    uint    `gorm:"primaryKey;column:id_user" json:"id_user"`
	QTY       uint    `gorm:"column:qty" json:"qty"`
	Product   Product `gorm:"foreignKey:id_product;references:id_product" json:"detail_product"`
	User      User    `gorm:"foreignKey:id_user;references:id_user" json:"detail_user"`
}

func (ShopCart) TableName() string {
	return "tbl_shoppingcart"
}

type PostCartRequest struct {
	IDProduct uint `json:"product" binding:"required"`
}

type UpdateCartRequest struct {
	IDProduct uint   `json:"product" binding:"required"`
	QTY       uint   `json:"qty" binding:"required"`
	Method    string `json:"method" binding:"required"`
}
