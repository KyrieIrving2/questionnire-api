package controllers

import (
	"github.com/gin-gonic/gin"
	"questionnaire_api/api/ctrl/assember"
	"questionnaire_api/api/ctrl/contract"
	"questionnaire_api/api/ctrl/contract/template"
	"questionnaire_api/api/ioc"
	"questionnaire_api/api/response"
	"questionnaire_api/api/user"
)

// @Summary 添加模板
// @Description 添加模板
// @Tags template
// @Accept json
// @Produce json
// @param Authorization header string true " "
// @param input body template.TemplateDto true "body"
// @Success 200 "ok"
// @Failure 400 "fail"
// @Router /template [post]
func AddTemplate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		input := template.TemplateDto{}

		if isValid, msg := input.Validate(ctx); !isValid {
			response.BadRequestError(msg, ctx)
			return
		}

		userName := user.GetUser(ctx)

		templateEntity := assember.ToTemplateEntity(&input)
		templateEntity.CreatedBy = userName.(string)
		templateEntity.UpdatedBy = userName.(string)

		id, err := ioc.DIContainer.TemplateSrv.Insert(templateEntity)
		if err != nil {
			response.BadRequestError(err.Error(), ctx)
			return
		}

		response.Ok(id, ctx)
	}
}

// @Summary 更新模板
// @Description 更新模板
// @Tags template
// @Accept json
// @Produce json
// @param Authorization header string true " "
// @param input body template.TemplateUpdateDto true "body"
// @Success 200 "ok"
// @Failure 400 "fail"
// @Router /template [put]
func UpdateTemplate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		input := template.TemplateUpdateDto{}

		if isValid, msg := input.Validate(ctx); !isValid {
			response.BadRequestError(msg, ctx)
			return
		}

		userName := user.GetUser(ctx)

		templateEntity := assember.ToTemplateUpdateEntity(&input)
		templateEntity.UpdatedBy = userName.(string)

		err := ioc.DIContainer.TemplateSrv.Update(templateEntity)

		if err != nil {
			response.BadRequestError(err.Error(), ctx)
			return
		}

		response.Ok("ok", ctx)
	}
}

// @Summary 模板列表
// @Description template list
// @Tags template
// @Accept  json
// @Produce json
// @Param name query string false " "
// @Param limit query int false " "
// @Param offset query int false " "
// @Success 200  "ok"
// @Router /template [get]
// +get /template
func SearchTemplate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		input := &template.TemplateFilterDto{}

		if isValid, msg := input.Validate(ctx); !isValid {
			response.BadRequestError(msg, ctx)
			return
		}

		templateFilterDo := assember.ToTemplateFilterDo(input)
		total, result := ioc.DIContainer.TemplateSrv.Search(templateFilterDo)
		pageVo := contract.PageVo{
			Total:   total,
			Results: result,
		}

		response.Ok(pageVo, ctx)
	}
}


// @Summary 删除模板
// @Description delete template
// @Tags template
// @Accept  json
// @Produce json
// @Param id path int false " "
// @Success 200  "ok"
// @Router /template/{id} [get]
// +get /template
func DeleteById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		response.Ok("ok",ctx)
	}
}
