package template

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"text/template"
)

func ApplyJson(in map[string]io.Reader, name string, param io.Reader) (string, error) {
	var data interface{}
	if param != nil {
		if err := json.NewDecoder(param).Decode(&data); err != nil {
			return "", fmt.Errorf("Cannot unmarshall json file : %w\n", err)
		}
	}

	var main *template.Template
	for n, r := range in {
		var tmpl *template.Template
		templateContent, err := ioutil.ReadAll(r)
		if err != nil {
			return "", fmt.Errorf("Cannot read input of %s", n)
		}
		if main == nil {
			main = template.New(n)
		}
		if n == main.Name() {
			tmpl = main
		} else {
			tmpl = main.New(n)
		}
		_, err = tmpl.Parse(string(templateContent))
		if err != nil {
			return "", err
		}
	}

	buf := bytes.Buffer{}
	err := main.ExecuteTemplate(&buf, name, data)
	return buf.String(), err
}
