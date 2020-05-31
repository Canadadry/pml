package cmd

import (
	"encoding/json"
	"fmt"
	"pml/pkg/template"
)

func Template(input string, param []byte) error {

	var dat interface{}
	if err := json.Unmarshal(param, &dat); err != nil {
		return fmt.Errorf("Cannot unmarshall json file : %w\n", err)
	}
	out, err := template.Apply(input, dat)
	if err != nil {
		return fmt.Errorf("failed to transform template : %w\n", err)
	}
	fmt.Println(out)
	return nil
}
