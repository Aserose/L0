package htmlPage

import (
	"html/template"
	"net/http"
	"path/filepath"
	"runtime"
)

type Html struct{}

func NewHtml() Html {
	return Html{}
}

func (h Html) TemplateExecute(w http.ResponseWriter, htmlFilename string) {
	template.Must(template.ParseFiles(h.GetFilepath(htmlFilename))).Execute(w, nil)
}

func (h Html) GetFilepath(name string) string {
	return filepath.Join(h.getDirectory(h.getFilename(), 0), name)
}

func (h Html) getDirectory(filename string, depth int) string {
	if depth != 0 {
		return h.getDirectory(filepath.Dir(filename), depth-1)
	}
	return filepath.Dir(filename)
}

func (h Html) getFilename() string {
	_, filename, _, _ := runtime.Caller(0)
	return filename
}
