package model

type TemplateFilterDo struct {
	Name   string `json:"name"`
	Limit  int64  `json:"limit"`
	Offset int64  `json:"offset"`
}
