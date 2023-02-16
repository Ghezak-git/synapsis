package models

import "time"

type TransactionHeader struct {
	ID                 uint               `gorm:"primaryKey;autoIncrement;column:id_transaction_header" json:"id_transaction_header"`
	NamaCustomer       string             `gorm:"type:varchar(150);not null;column:name_customer" json:"name_customer"`
	NoTransaction      string             `gorm:"type:varchar(150);column:no_transaction" json:"no_transaction"`
	PaymentMethod      string             `gorm:"type:varchar(150);column:payment_method_name" json:"payment_method_name"`
	PaymentDescription string             `gorm:"column:payment_method_desc" json:"payment_method_desc"`
	Total              uint               `gorm:"column:total" json:"total"`
	Deleted            string             `gorm:"type:varchar(10);column:deleted;default:'N'" json:"deleted"`
	IDCreated          uint               `gorm:"column:id_created" json:"id_created"`
	IDUpdated          uint               `gorm:"column:id_updated" json:"id_updated"`
	CreatedAt          time.Time          `gorm:"column:created_at" json:"created_at"`
	UpdatedAt          time.Time          `gorm:"column:updated_at" json:"updated_at"`
	Details            []TransactioDetail `gorm:"foreignKey:id_transaction_header;references:id_transaction_header" json:"detail"`
}

func (TransactionHeader) TableName() string {
	return "tbl_transaction_header"
}

type TransactioDetail struct {
	ID                 uint      `gorm:"primaryKey;autoIncrement;column:id_transaction_detail" json:"id_transaction_detail"`
	IDHeader           uint      `gorm:"column:id_transaction_header" json:"id_transaction_header"`
	NamaProduct        string    `gorm:"type:varchar(150);column:name_product" json:"name_product"`
	DescriptionProduct string    `gorm:"column:description_product" json:"description_product"`
	QTY                uint      `gorm:"column:qty" json:"qty"`
	Price              uint      `gorm:"column:price_product" json:"price_product"`
	Total              uint      `gorm:"column:total" json:"total"`
	Deleted            string    `gorm:"type:varchar(10);column:deleted;default:'N'" json:"deleted"`
	IDCreated          uint      `gorm:"column:id_created" json:"id_created"`
	IDUpdated          uint      `gorm:"column:id_updated" json:"id_updated"`
	CreatedAt          time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt          time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (TransactioDetail) TableName() string {
	return "tbl_transaction_detail"
}

type PaymentMethod struct {
	ID          uint      `gorm:"primaryKey;column:iIdPaymentMethod" json:"iIdPaymentMethod"`
	Nama        string    `gorm:"type:varchar(150);column:name" json:"name"`
	Description string    `gorm:"column:description" json:"description"`
	Deleted     string    `gorm:"type:varchar(10);column:deleted;default:'N'" json:"deleted"`
	IDCreated   uint      `gorm:"column:id_created" json:"id_created"`
	IDUpdated   uint      `gorm:"column:id_updated" json:"id_updated"`
	CreatedAt   time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (PaymentMethod) TableName() string {
	return "master_payment_method"
}

type TransationRequest struct {
	Page       string `json:"page,omitempty" `
	PerPage    string `json:"perPage,omitempty"`
	Search     string `json:"search,omitempty"`
	Sort       string `json:"sort,omitempty"`
	SortColumn string `json:"sortColumn,omitempty"`
}

type TrasactionPostHeaderRequest struct {
	PaymentMethod      string                        `json:"payment_method_name,omitempty" binding:"required"`
	PaymentDescription string                        `json:"payment_method_desc,omitempty" binding:"required"`
	Total              uint                          `json:"total,omitempty" binding:"required"`
	Detail             []TrasactionPostDetailRequest `json:"detail,omitempty" binding:"required"`
}

type TrasactionPostDetailRequest struct {
	NamaProduct        string `json:"name_product,omitempty" binding:"required"`
	DescriptionProduct string `json:"description_product,omitempty" binding:"required"`
	Price              uint   `json:"price_product,omitempty" binding:"required"`
	QTY                uint   `json:"qty,omitempty" binding:"required"`
}
