package persistence

import (
	"questionnaire_api/domain/template/entity"
	"questionnaire_api/domain/template/repository/facade"
	"questionnaire_api/orm"
)

type TemplateRepository struct {
}

var _ facade.ITemplateRepository = new(TemplateRepository)

func (t *TemplateRepository) Insert(template *entity.Template) (int, error) {
	db := orm.GetInstance().GetMysqlDB()

	if err := db.Create(template).Error; err != nil {
		return 0, err
	}

	return template.ID, nil
}

func (t *TemplateRepository) Update(template *entity.Template) error {
	db := orm.GetInstance().GetMysqlDB()

	return db.Model(template).Updates(
		map[string]interface{}{
			"name":       template.Name,
			"updated_at": template.UpdatedAt,
			"updated_by": template.UpdatedBy}).Error
}
