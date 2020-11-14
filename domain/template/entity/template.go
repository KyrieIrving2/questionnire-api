package entity

import (
	"time"
)

// Template table comment
type Template struct {
	// CreatedAt column comments
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	// CreatedBy column comments
	CreatedBy string `gorm:"column:created_by" json:"created_by"`
	// ID column comments
	ID int `gorm:"primary_key;column:id" json:"id"`
	// IsCurrent column comments
	IsCurrent int `gorm:"column:is_current" json:"is_current"`
	// Name column comments
	Name string `gorm:"column:name" json:"name"`
	// UpdatedAt column comments
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
	// UpdatedBy column comments
	UpdatedBy string `gorm:"column:updated_by" json:"updated_by"`
}

// TableName sets the insert table name for this struct type
func (t *Template) TableName() string {
	return "template"
}
