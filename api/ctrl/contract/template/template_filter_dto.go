package template

import "github.com/gin-gonic/gin"

type TemplateFilterDto struct {
	Name   string `json:"name" form:"name"`
	Limit  int64  `json:"limit" form:"limit"`
	Offset int64  `json:"offset" form:"offset"`
}

//校验参数
func (t *TemplateFilterDto) Validate(ctx *gin.Context) (bool, string) {
	err := ctx.ShouldBind(t)
	if err != nil {
		return false, "请求参数解析异常,请核对"
	}

	return true, ""
}