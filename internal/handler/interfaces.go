package handler

type TemplateBuilder interface {
	Build(templateFile string, templateData any) ([]byte, error)
}
