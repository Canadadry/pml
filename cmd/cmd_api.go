package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func Api(pmlFolder string) error {
	pmls, err := list(pmlFolder, ".pml")
	if err != nil {
		return err
	}
	log.Println("starting server on " + pmlFolder)
	http.HandleFunc("/", HelloServer(pmlFolder, pmls))
	return http.ListenAndServe(":8080", nil)
}

func HelloServer(pmlFolder string, pmls []string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Path[1:]

		if len(id) > 0 {
			for _, pml := range pmls {
				if pml == id {
					log.Printf("rendering %s\n", pmlFolder+id+".pml")
					file, err := ioutil.ReadFile(pmlFolder + id + ".pml")
					if err != nil {
						log.Println("Cannot find file " + pmlFolder + id + ".pml")
						fmt.Fprintf(w, "Cannot find file "+pmlFolder+id+".pml")
						return
					}

					bodyBytes, err := ioutil.ReadAll(r.Body)
					if err != nil {
						log.Printf("Cannot read request : %v", err)
						fmt.Fprintf(w, "Cannot read request : %v", err)
						return
					}
					w.Header().Set("Content-Type", "application/pdf")
					w.WriteHeader(http.StatusOK)
					err = Full(string(file), w, bodyBytes)
					if err != nil {
						log.Println("error writing result", err)
						fmt.Fprintf(w, "error writing result : %v", err)
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
