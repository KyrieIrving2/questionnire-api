package persistence

import (
	"questionnaire_api/domain/template/entity"
	"questionnaire_api/domain/template/model"
	"questionnaire_api/domain/template/repository/facade"
	"questionnaire_api/orm"
)

type TemplateQuery struct {
}

var _ facade.ITemplateQuery = new(TemplateQuery)

func (t *TemplateQuery) Search(templateFilter *model.TemplateFilterDo) (total int, results []*entity.Template) {
	db := orm.GetInstance().GetMysqlDB()
	db.Where("name = ?", templateFilter.Name).
		Count(&total).
		Order("id asc").
		Limit(templateFilter.Limit).
		Offset(templateFilter.Offset).
		Find(&results)

	return
}
