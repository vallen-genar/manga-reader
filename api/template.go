package api

import (
	"html/template"
	"io/ioutil"
	"log"
	"os"
)

const basePath = "template"

func populateTemplates() map[string]*template.Template {
	result := make(map[string]*template.Template)

	funcMap := template.FuncMap{
		// The name "inc" is what the function will be called in the template text.
		"even": func(i int) bool {
			return i%2 == 0
		},
		"add": func(a, b int) int {
			return a + b
		},
		"more": func(a, b int) bool {
			return a > b
		},
		"mod": func(a, b int) bool {
			return a%b == 0
		},
	}

	layout := template.Must(template.ParseFiles(basePath + "/_layout.html"))
	template.Must(layout.ParseFiles(basePath+"/_header.html", basePath+"/_footer.html"))
	dir, err := os.Open(basePath + "/content")
	if err != nil {
		panic("Failed to open template blocks directory: " + err.Error())
	}

	fis, err := dir.Readdir(-1)
	if err != nil {
		panic("Failed to read contents of content directory: " + err.Error())
	}
	for _, fi := range fis {
		f, err := os.Open(basePath + "/content/" + fi.Name())
		if err != nil {
			panic("Failed to open template '" + fi.Name() + "'")
		}
		content, err := ioutil.ReadAll(f)
		if err != nil {
			panic("Failed to read content from file '" + fi.Name() + "'")
		}
		f.Close()
		tmpl := template.Must(layout.Clone()).Funcs(funcMap)
		_, err = tmpl.Parse(string(content))
		if err != nil {
			log.Println(err.Error())
			panic("Failed to parse contents of '" + fi.Name() + "' as template")
		}
		result[fi.Name()] = tmpl
	}
	return result
}