package main

import (
	"flag"
	"fmt"
	"github.com/canadadry/pml/compiler"
	"github.com/canadadry/pml/pkg/eval"
	"github.com/canadadry/pml/pkg/i18n"
	"github.com/canadadry/pml/pkg/pdf"
	"github.com/canadadry/pml/pkg/svg"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {

	var (
		filename  string
		main      string
		paramfile string
		output    string
		transfile string
		local     string
	)

	flag.StringVar(&filename, "in", "", "entry pml filename or directory (will only load sub .pml files)")
	flag.StringVar(&main, "main", "", "if in is a folder, main should contain the main filename")
	flag.StringVar(&paramfile, "param", "", "param for pml filename")
	flag.StringVar(&output, "out", "out.pdf", "pdf output for renderer mode")
	flag.StringVar(&transfile, "trans", "", "specify translation file if needed")
	flag.StringVar(&local, "local", "fr", "which local of the trans file to load")

	flag.Parse()

	if len(filename) == 0 {
		flag.PrintDefaults()
		return nil
	}

	files, err := LoadInputFiles(filename)
	if err != nil {
		return err
	}
	defer CloseFiles(files)

	if len(files) == 1 {
		for n := range files {
			main = n
		}
	} else {
		if len(main) == 0 {
			return fmt.Errorf("You must select your main file")
		}
	}

	var fParam io.ReadCloser
	fParam, err = os.Open(paramfile)
	if err != nil {
		fParam = nil
	} else {
		defer fParam.Close()
	}

	fOut, err := os.Create(output)
	if err != nil {
		return fmt.Errorf("Cannot create output file " + output)
	}
	defer fOut.Close()

	var fTranslation io.ReadCloser
	fTranslation, err = os.Open(transfile)
	if err != nil {
		fTranslation = nil
	} else {
		defer fTranslation.Close()
	}

	tr, err := i18n.LoadFromCsv(fTranslation, local)
	if err != nil {
		return fmt.Errorf("Cannot read tranlation file %v", err)
	}

	env := compiler.Env{
		Input:  files,
		Main:   main,
		Output: fOut,
		Param:  fParam,
		Pdf:    pdf.New(svg.Draw),
		Funcs: map[string]interface{}{
			"eval": templateEval,
			"data": BuildDataMap,
			"tr":   Translate(tr),
			"upper": func(v interface{}) string {
				return strings.ToUpper(fmt.Sprintf("%v", v))
			},
		},
	}

	return compiler.Run(env)
}

func LoadInputFiles(filename string) (map[string]io.Reader, error) {
	files := map[string]io.Reader{}

	f, err := os.Open(filename)
	if err != nil {
		return files, fmt.Errorf("Cannot find file '%s' : %w", filename, err)
	}
	stat, err := f.Stat()
	if err != nil {
		return files, fmt.Errorf("Cannot stat file '%s' : %w", filename, err)
	}

	if !stat.IsDir() {
		files[filename] = f
		return files, nil
	}

	paths, err := glob(filename, ".pml")
	if err != nil {
		return files, fmt.Errorf("Cannot glob '%s' : %w", filename, err)
	}

	for _, p := range paths {
		f, err := os.Open(p)
		if err != nil {
			return files, fmt.Errorf("Cannot find file '%s' : %w", p, err)
		}
		stat, err := f.Stat()
		if err != nil {
			return files, fmt.Errorf("Cannot stat file '%s' : %w", p, err)
		}
		if stat.IsDir() {
			continue
		}
		files[p] = f
	}

	return files, nil
}

func glob(dir string, ext string) ([]string, error) {

	files := []string{}
	err := filepath.Walk(dir, func(path string, f os.FileInfo, err error) error {
		if filepath.Ext(path) == ext {
			files = append(files, path)
		}
		return nil
	})

	return files, err
}

func templateEval(values ...interface{}) string {
	str := ""
	for _, v := range values {
		str = fmt.Sprintf("%s %v", str, v)
	}
	result, err := eval.Eval(str)
	if err != nil {
		return err.Error()
	}
	return fmt.Sprintf("%v", result)
}

func BuildDataMap(values ...interface{}) (map[string]interface{}, error) {
	ret := map[string]interface{}{}
	if len(values)%2 == 1 {
		return nil, fmt.Errorf("Must have a pair number of param : a key and a value")
	}
	for i := 0; i < len(values); i += 2 {
		ret[fmt.Sprintf("%v", values[i+0])] = values[i+1]
	}
	return ret, nil
}

func Translate(tr i18n.Translation) func(string, ...interface{}) (string, error) {
	return func(key string, p ...interface{}) (string, error) {
		if len(p)%2 == 1 {
			return "", fmt.Errorf("Must have a pair number of param : a key and a value")
		}
		ps := []i18n.Param{}
		for i := 0; i < len(p); i += 2 {
			ps = append(ps, i18n.Param{
				Old: fmt.Sprintf("%v", p[i+0]),
				New: fmt.Sprintf("%v", p[i+1]),
			})
		}
		return tr.Trans(key, ps), nil
	}
}

func CloseFiles(files map[string]io.Reader) {
	for _, f := range files {
		fcloser, ok := f.(io.ReadCloser)
		if !ok {
			continue
		}
		fcloser.Close()
	}
}
