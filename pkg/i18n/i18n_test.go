package i18n

import (
	"bytes"
	"reflect"
	"testing"
)

func TestLoadFromCsv(t *testing.T) {
	tests := []struct {
		in       string
		local    string
		expected Translation
	}{
		{
			in: `,fr,en,de,it,es,nl,pt,ls
appendices.title,Annexes,Appendices,Anhänge,Appendici,Apéndices,Bijlagen,Appêndices,ls
authorization.applicant_signature_title,Signature du candidat*,Applicant signature*,Unterschrift des Antragstellers*,Apporre la firma*,Firma del solicitante*,Handtekening kandidaat*,Assinatura do candidato*,ls
`,
			local: "fr",
			expected: Translation{
				"appendices.title":                        "Annexes",
				"authorization.applicant_signature_title": "Signature du candidat*",
			},
		},
		{
			in: `,fr,en,de,it,es,nl,pt,ls
appendices.title,Annexes,Appendices,Anhänge,Appendici,Apéndices,Bijlagen,Appêndices,ls
authorization.applicant_signature_title,Signature du candidat*,Applicant signature*,Unterschrift des Antragstellers*,Apporre la firma*,Firma del solicitante*,Handtekening kandidaat*,Assinatura do candidato*,ls
`,
			local: "pt",
			expected: Translation{
				"appendices.title":                        "Appêndices",
				"authorization.applicant_signature_title": "Assinatura do candidato*",
			},
		},
	}

	for i, tt := range tests {
		result, err := LoadFromCsv(bytes.NewBufferString(tt.in), tt.local)
		if err != nil {
			t.Fatalf("[%d] failed %v", i, err)
		}
		if !reflect.DeepEqual(result, tt.expected) {
			t.Fatalf("[%d] exp \n%#v\n got \n%#v\n", i, tt.expected, result)
		}
	}
}

func TestTrans(t *testing.T) {
	tests := []struct {
		in       Translation
		key      string
		params   []Param
		expected string
	}{
		{
			in:       Translation{},
			key:      "fake",
			params:   nil,
			expected: "",
		},
		{
			in:       Translation{"key": "value"},
			key:      "key",
			params:   nil,
			expected: "value",
		},
		{
			in:  Translation{"key": "valuelu"},
			key: "key",
			params: []Param{
				{Old: "lu", New: "test"},
			},
			expected: "vatestetest",
		},
		{
			in:  Translation{"key": "valuelu"},
			key: "key",
			params: []Param{
				{Old: "lu", New: "test"},
				{Old: "stet", New: ""},
			},
			expected: "vateest",
		},
	}

	for i, tt := range tests {
		result := tt.in.Trans(tt.key, tt.params)
		if result != tt.expected {
			t.Fatalf("[%d] exp \n%#v\n got \n%#v\n", i, tt.expected, result)
		}
	}
}
