package files

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

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

func CloseFiles(files map[string]io.Reader) {
	for _, f := range files {
		fcloser, ok := f.(io.ReadCloser)
		if !ok {
			continue
		}
		fcloser.Close()
	}
}
