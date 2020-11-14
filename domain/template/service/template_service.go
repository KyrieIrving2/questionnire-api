package service

import (
	"questionnaire_api/api/ctrl/assember"
	"questionnaire_api/api/ctrl/contract/template"
	models "questionnaire_api/domain/template/entity"
	"questionnaire_api/domain/template/model"
	"questionnaire_api/domain/template/repository/facade"
)

type TemplateSrv struct {
	TemplateRepository facade.ITemplateRepository `inject:"TemplateRepository"`
	TemplateQuery      facade.ITemplateQuery      `inject:"TemplateQuery"`
}

var _ ITemplateSrv = new(TemplateSrv)

func (t *TemplateSrv) Insert(template *models.Template) (int, error) {
	return t.TemplateRepository.Insert(template)
}

func (t *TemplateSrv) Update(template *models.Template) error {
	return t.TemplateRepository.Update(template)
}

func (t *TemplateSrv) Search(templateFilter *model.TemplateFilterDo) (total int, results []*template.TemplateVo) {
	total, templates := t.TemplateQuery.Search(templateFilter)
	results = assember.ToTemplateArrVo(templates)
	return
}
