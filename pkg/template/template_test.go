package template

import (
	"bytes"
	"io"
	"testing"
)

func TestApplyJson(t *testing.T) {
	tests := []struct {
		in     map[string]string
		main   string
		data   string
		result string
	}{
		{
			in: map[string]string{
				"test": "test",
			},
			main:   "test",
			data:   "{}",
			result: "test",
		},
		{
			in: map[string]string{
				"test": "hello {{ .value}}",
			},
			main:   "test",
			data:   `{"value":"world"}`,
			result: "hello world",
		},
		{
			in: map[string]string{
				"foo": `foo {{ template "bar" }}`,
				"bar": "bar",
			},
			main:   "foo",
			data:   "{}",
			result: "foo bar",
		},
		{
			in: map[string]string{
				"foo": `foo {{ template "bar" .}}`,
				"bar": "bar {{.value}}",
			},
			main:   "foo",
			data:   `{"value":"world"}`,
			result: "foo bar world",
		},
	}
	for i, tt := range tests {
		in := map[string]io.Reader{}
		for n, s := range tt.in {
			in[n] = bytes.NewBufferString(s)
		}
		out, err := ApplyJson(
			in,
			tt.main,
			bytes.NewBufferString(tt.data),
		)
		if err != nil {
			t.Fatalf("[%d] failed %v", i, err)
		}
		if tt.result != out {
			t.Fatalf("[%d] failed \n%s\n%s", i, tt.result, out)
		}
	}
}
