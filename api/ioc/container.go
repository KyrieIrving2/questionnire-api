package ioc

import (
	"github.com/facebookgo/inject"
	"questionnaire_api/domain/template/repository/persistence"
	"questionnaire_api/domain/template/service"
)

// DIContainer
var DIContainer Container

// Container
type Container struct {
	TemplateSrv service.ITemplateSrv `inject:"TemplateSrv"`
}

// ioc
func InitIoc() {
	var g inject.Graph
	handleErr(g.Provide(&inject.Object{Value: &DIContainer}))
	handle(&g)
	handleErr(g.Populate())
}

// handle
func handle(g *inject.Graph) {
	handleErr(g.Provide(
		&inject.Object{Value: &service.TemplateSrv{}, Name: "TemplateSrv"},
		&inject.Object{Value: &persistence.TemplateRepository{}, Name: "TemplateRepository"},
		&inject.Object{Value: &persistence.TemplateQuery{}, Name: "TemplateQuery"},
	))
}

// handle err
func handleErr(err error) {
	if err != nil {
		panic(err)
	}
}
