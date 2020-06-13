package template

import (
	"bytes"
	"text/template"
)

func Apply(templateContent string, param interface{}) (string, error) {

	pmlTemplate, err := template.New("pml").Parse(templateContent)
	if err != nil {
		return "", err
	}

	s := ""
	buf := bytes.NewBufferString(s)

	err = pmlTemplate.Execute(buf, param)
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}
