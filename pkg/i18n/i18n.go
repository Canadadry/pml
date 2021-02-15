package i18n

import (
	"encoding/csv"
	// "fmt"
	"io"
	"strings"
)

type Translation map[string]string

type Param struct {
	Old string
	New string
}

func LoadFromCsv(file io.Reader, local string) (Translation, error) {
	if file == nil {
		return Translation{}, nil
	}
	csvReader := csv.NewReader(file)
	csvReader.Comma = ','

	records, err := csvReader.ReadAll()
	if err != nil {
		return nil, err
	}
	if len(records) < 2 {
		return Translation{}, nil
	}

	localIndex := 0
	for i, l := range records[0] {
		if l == local {
			localIndex = i
		}
	}

	t := Translation{}
	for _, line := range records[1:] {
		t[line[0]] = line[localIndex]
		// fmt.Println(line[0], ":", line[localIndex])
	}
	return t, nil
}

func (t Translation) Trans(key string, params []Param) string {
	str, ok := t[key]
	if !ok {
		return ""
	}
	for _, p := range params {
		str = strings.ReplaceAll(str, p.Old, p.New)
	}
	return str
}
