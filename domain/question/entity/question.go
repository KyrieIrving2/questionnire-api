package entity

import (
	"time"
)

// Question table comment
type Question struct {
	// CreatedAt column comments
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	// CreatedBy column comments
	CreatedBy string `gorm:"column:created_by" json:"created_by"`
	// ID column comments
	ID int `gorm:"primary_key;column:id" json:"id"`
	// Name column comments
	Name string `gorm:"column:name" json:"name"`
	// Priority column comments
	Priority int `gorm:"column:priority" json:"priority"`
	// Status column comments
	Status int `gorm:"column:status" json:"status"`
	// TemplateID column comments
	TemplateID int `gorm:"column:template_id" json:"template_id"`
	// TypeID column comments
	TypeID int `gorm:"column:type_id" json:"type_id"`
	// UpdatedAt column comments
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
	// UpdatedBy column comments
	UpdatedBy string `gorm:"column:updated_by" json:"updated_by"`
}

// TableName sets the insert table name for this struct type
func (q *Question) TableName() string {
	return "question"
}
