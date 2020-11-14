package service

import (
	"questionnaire_api/api/ctrl/contract/template"
	"questionnaire_api/domain/template/entity"
	"questionnaire_api/domain/template/model"
)

type ITemplateSrv interface {
	Insert(template *entity.Template) (int, error)
	Update(template *entity.Template) error
	Search(templateFilter *model.TemplateFilterDo) (total int, results []*template.TemplateVo)
}
