package assember

import (
	"questionnaire_api/api/ctrl/contract/template"
	"questionnaire_api/domain/template/entity"
	"questionnaire_api/domain/template/model"
	"time"
)

func ToTemplateEntity(dto *template.TemplateDto) *entity.Template {
	return &entity.Template{
		Name:      dto.Name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func ToTemplateUpdateEntity(dto *template.TemplateUpdateDto) *entity.Template {
	return &entity.Template{
		ID:        dto.Id,
		Name:      dto.Name,
		UpdatedAt: time.Now(),
	}
}

func ToTemplateFilterDo(dto *template.TemplateFilterDto) *model.TemplateFilterDo {
	return &model.TemplateFilterDo{
		Name:   dto.Name,
		Limit:  dto.Limit,
		Offset: dto.Offset,
	}
}

func ToTemplateArrVo(templates []*entity.Template) []*template.TemplateVo {
	var templateArrVo []*template.TemplateVo

	for _, v := range templates {
		templateArrVo = append(templateArrVo, &template.TemplateVo{
			CreatedAt: v.CreatedAt,
			CreatedBy: v.CreatedBy,
			ID:        v.ID,
			IsCurrent: v.IsCurrent,
			Name:      v.Name,
			UpdatedAt: v.UpdatedAt,
			UpdatedBy: v.UpdatedBy,
		})
	}

	return templateArrVo
}
