package entity

import (
	"github.com/guregu/null"
)

// LookupResource table comment
type LookupResource struct {
	// ID column comments
	ID int `gorm:"column:id" json:"id"`
	// Priority column comments
	Priority null.Int `gorm:"column:priority" json:"priority"`
	// TypeCode column comments
	TypeCode null.String `gorm:"column:type_code" json:"type_code"`
	// TypeName column comments
	TypeName null.String `gorm:"column:type_name" json:"type_name"`
}

// TableName sets the insert table name for this struct type
func (l *LookupResource) TableName() string {
	return "lookup_resource"
}
