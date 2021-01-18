package template

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"text/template"
)

func Apply(in io.Reader, param io.Reader) (string, error) {
	templateContent, err := ioutil.ReadAll(in)
	if err != nil {
		return "", fmt.Errorf("Cannot read input")
	}

	var data interface{}
	err = json.NewDecoder(param).Decode(&data)
	if err != nil {
		return "", fmt.Errorf("Cannot unmarshall json file : %w\n", err)
	}

	pmlTemplate, err := template.New("pml").Parse(string(templateContent))
	if err != nil {
		return "", err
	}

	s := ""
	buf := bytes.NewBufferString(s)

	err = pmlTemplate.Execute(buf, param)
	return buf.String(), err
}
