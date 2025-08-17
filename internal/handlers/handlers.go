package handlers

import (
	"html/template"
	"log"
	"os"
	"path/filepath"

	"github.com/gorilla/mux"
)

// Just sets the routes and displays html
type DisplayHandler interface{
	Routes(router *mux.Router)
}

var handlerRegistry []DisplayHandler

func RegisterHandler(h DisplayHandler) {
	handlerRegistry = append(handlerRegistry, h)
}
func GetHandlers() []DisplayHandler {
	return handlerRegistry
}

// Takes './internal/handlers' as base-path.
// Keep in mind that paths[0] must be the base/root-template
// that uses all other templates! 
func LoadTemplates(paths []string) *template.Template{
	wd, err := os.Getwd()
  if err != nil{
    log.Fatal("couldn't get working directory: ", err)
  }
	base := filepath.Join(wd, "internal", "handlers")
	var full = make([]string, len(paths))
	for i, p := range paths{
		full[i] = filepath.Join(base,p)
	}
	funcMap := template.FuncMap{
		"arr": func (item ...any) []any { return item },
	}
	// add funcMap to base-template
	first := filepath.Base(full[0])
	tmpl := template.New(first).Funcs(template.FuncMap(funcMap))
	tmpl = template.Must(tmpl.ParseFiles(full...))
	return tmpl
}

