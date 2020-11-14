package entity

import (
	"github.com/guregu/null"
)

// Answer table comment
type Answer struct {
	// ID column comments
	ID int `gorm:"primary_key;column:id" json:"id"`
	// QuestionID column comments
	QuestionID int `gorm:"column:question_id" json:"question_id"`
	// QuestionItemID column comments
	QuestionItemID null.Int `gorm:"column:question_item_id" json:"question_item_id"`
	// UserID column comments
	UserID int `gorm:"column:user_id" json:"user_id"`
	// Value column comments
	Value null.String `gorm:"column:value" json:"value"`
}

// TableName sets the insert table name for this struct type
func (a *Answer) TableName() string {
	return "answer"
}
