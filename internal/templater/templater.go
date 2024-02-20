package templater

import (
	"bytes"
	"fmt"
	"html/template"
)

type Templater struct {
	templatesDir string
}

func New(templatesDir string) *Templater {
	return &Templater{
		templatesDir: templatesDir,
	}
}

func (t *Templater) Build(templateFile string, data any) ([]byte, error) {
	tmpl, err := template.New(templateFile).ParseFiles(t.templatesDir + "/" + templateFile)
	if err != nil {
		fmt.Println(err.Error())
		return []byte(err.Error()), err
	}

	body := new(bytes.Buffer)
	err = tmpl.Execute(body, data)
	if err != nil {
		return []byte(err.Error()), err
	}

	return body.Bytes(), nil
}
