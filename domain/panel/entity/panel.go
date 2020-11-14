package entity


// Panel table comment
type Panel struct {
	// ID column comments
	ID int `gorm:"primary_key;column:id" json:"id"`
	// TemplateID column comments
	TemplateID int `gorm:"column:template_id" json:"template_id"`
	// TypeID column comments
	TypeID int `gorm:"column:type_id" json:"type_id"`
	// Value column comments
	Value string `gorm:"column:value" json:"value"`
}

// TableName sets the insert table name for this struct type
func (p *Panel) TableName() string {
	return "panel"
}
