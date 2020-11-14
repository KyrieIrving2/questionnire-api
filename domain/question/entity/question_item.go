package entity


// QuestionItem table comment
type QuestionItem struct {
	// ID column comments
	ID int `gorm:"primary_key;column:id" json:"id"`
	// Name column comments
	Name string `gorm:"column:name" json:"name"`
	// Priority column comments
	Priority int `gorm:"column:priority" json:"priority"`
	// QuestionID column comments
	QuestionID int `gorm:"column:question_id" json:"question_id"`
}

// TableName sets the insert table name for this struct type
func (q *QuestionItem) TableName() string {
	return "question_item"
}
