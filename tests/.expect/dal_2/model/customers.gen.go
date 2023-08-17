// Code generated by github.com/dieagenturverwaltung/gorm-gen. DO NOT EDIT.
// Code generated by github.com/dieagenturverwaltung/gorm-gen. DO NOT EDIT.
// Code generated by github.com/dieagenturverwaltung/gorm-gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameCustomer = "customers"

// Customer mapped from table <customers>
type Customer struct {
	ID        int64          `gorm:"column:id;primaryKey;autoIncrement:true" json:"-"`
	CreatedAt *time.Time     `gorm:"column:created_at" json:"-"`
	UpdatedAt *time.Time     `gorm:"column:updated_at" json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;index:idx_customers_deleted_at,priority:1" json:"-"`
	BankID    *int64         `gorm:"column:bank_id" json:"-"`
}

// TableName Customer's table name
func (*Customer) TableName() string {
	return TableNameCustomer
}
