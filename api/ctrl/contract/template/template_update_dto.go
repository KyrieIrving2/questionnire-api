package template

import "github.com/gin-gonic/gin"

type TemplateUpdateDto struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func (t *TemplateUpdateDto) Validate(ctx *gin.Context) (bool, string) {
	err := ctx.ShouldBind(t)
	if err != nil {
		return false, "请求参数解析异常,请核对"
	}

	if t.Id <= 0 {
		return false, "请输入模板id"
	}
	if len(t.Name) <= 0 {
		return false, "请输入模板名"
	}

	return true, ""
}
