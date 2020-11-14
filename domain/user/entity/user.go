package entity

import (
	"github.com/guregu/null"
	"time"
)

// User table comment
type User struct {
	// CreatedAt column comments
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	// CreatedBy column comments
	CreatedBy string `gorm:"column:created_by" json:"created_by"`
	// Gender column comments
	Gender null.String `gorm:"column:gender" json:"gender"`
	// GenderID column comments
	GenderID null.Int `gorm:"column:gender_id" json:"gender_id"`
	// ID column comments
	ID int `gorm:"primary_key;column:id" json:"id"`
	// PositionID column comments
	PositionID null.Int `gorm:"column:position_id" json:"position_id"`
	// PositionName column comments
	PositionName null.String `gorm:"column:position_name" json:"position_name"`
	// RegionID column comments
	RegionID null.Int `gorm:"column:region_id" json:"region_id"`
	// RegionName column comments
	RegionName null.String `gorm:"column:region_name" json:"region_name"`
	// UpdatedAt column comments
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
	// UpdatedBy column comments
	UpdatedBy string `gorm:"column:updated_by" json:"updated_by"`
	// UserName column comments
	UserName string `gorm:"column:user_name" json:"user_name"`
}

// TableName sets the insert table name for this struct type
func (u *User) TableName() string {
	return "user"
}
