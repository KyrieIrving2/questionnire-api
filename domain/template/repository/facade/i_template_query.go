package facade

import (
	"questionnaire_api/domain/template/entity"
	"questionnaire_api/domain/template/model"
)

type ITemplateQuery interface {
	Search(templateFilter *model.TemplateFilterDo) (total int, results []*entity.Template)
}
