package template

import "github.com/gin-gonic/gin"

type TemplateDto struct {
	Name string `json:"name"`
}

func (t *TemplateDto) Validate(ctx *gin.Context) (bool, string) {
	err := ctx.ShouldBind(t)
	if err != nil {
		return false, "请求参数解析异常,请核对"
	}
	if len(t.Name) <= 0 {
		return false, "请输入模板名"
	}

	return true, ""
}
