package template

import "time"

type TemplateVo struct {
	CreatedAt time.Time `json:"createdAt"`
	CreatedBy string    `json:"createdBy"`
	ID        int       `json:"id"`
	IsCurrent int       `json:"isCurrent"`
	Name      string    `json:"name"`
	UpdatedAt time.Time `json:"updatedAt"`
	UpdatedBy string    `json:"updatedBy"`
}
