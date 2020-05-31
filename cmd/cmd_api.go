package cmd

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func Api(pmlFolder string) error {
	pmls, err := list(pmlFolder, ".pml")
	if err != nil {
		return err
	}
	fmt.Println("starting server on " + pmlFolder)
	http.HandleFunc("/", HelloServer(pmlFolder, pmls))
	return http.ListenAndServe(":8080", nil)
}

func HelloServer(pmlFolder string, pmls []string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Path[1:]

		if len(id) > 0 {
			for _, pml := range pmls {
				if pml == id {
					fmt.Printf("rendering %s\n", pmlFolder+id+".pml")
					file, err := ioutil.ReadFile(pmlFolder + id + ".pml")
					if err != nil {
						fmt.Println("Cannot find file " + pmlFolder + id + ".pml")
						fmt.Fprintf(w, "Cannot find file "+pmlFolder+id+".pml")
						return
					}
					w.Header().Set("Content-Type", "application/pdf")
					w.WriteHeader(http.StatusOK)
					err = Renderer(string(file), w)
					if err != nil {
						fmt.Println("error writing result", err)
					}
					return
				}
			}
		}

		fmt.Fprintf(w, "Hello, runnable files are : %v", pmls)
	}
}

func list(path string, ext string) ([]string, error) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}

	templates := []string{}

	for _, file := range files {
		filename := file.Name()
		if filename[len(filename)-len(ext):] == ext {
			templates = append(templates, filename[:len(filename)-len(ext)])
		}
	}

	return templates, nil
}
