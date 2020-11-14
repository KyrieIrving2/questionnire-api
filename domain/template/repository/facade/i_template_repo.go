package facade

import "questionnaire_api/domain/template/entity"

type ITemplateRepository interface {
	Insert(template *entity.Template) (int, error)
	Update(template *entity.Template) error
}
